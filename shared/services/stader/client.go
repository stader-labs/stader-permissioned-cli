/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.2.1]

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package stader

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/stader-labs/stader-node/shared"
	"io"
	"io/ioutil"
	"log"
	"math/big"
	"net"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/a8m/envsubst"
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh"

	"github.com/alessio/shellescape"
	"github.com/blang/semver/v4"
	externalip "github.com/glendc/go-external-ip"
	"github.com/mitchellh/go-homedir"

	"github.com/stader-labs/stader-node/shared/services/config"
	cfgtypes "github.com/stader-labs/stader-node/shared/types/config"
	staderUtils "github.com/stader-labs/stader-node/shared/utils/stdr"
)

// Config
const (
	InstallerURL = "https://staderlabs.com/eth/releases" + shared.BinaryBucket + "/%s/install.sh"

	SettingsFile             string = "user-settings.yml"
	BackupSettingsFile       string = "user-settings-backup.yml"
	PrometheusConfigTemplate string = "prometheus.tmpl"
	PrometheusFile           string = "prometheus.yml"

	APIContainerSuffix string = "_api"
	APIBinPath         string = "/go/bin/stader"

	templatesDir string = "templates"
	overrideDir  string = "override"
	runtimeDir   string = "runtime"

	templateSuffix    string = ".tmpl"
	composeFileSuffix string = ".yml"

	nethermindPruneStarterCommand string = "dotnet /setup/NethermindPruneStarter/NethermindPruneStarter.dll"
	nethermindAdminUrl            string = "http://127.0.0.1:7434"

	DebugColor = color.FgYellow
)

// Get the external IP address. Try finding an IPv4 address first to:
// * Improve peer discovery and node performance
// * Avoid unnecessary container restarts caused by switching between IPv4 and IPv6
func getExternalIP() (net.IP, error) {
	// Try IPv4 first
	ip4Consensus := externalip.DefaultConsensus(nil, nil)
	ip4Consensus.UseIPProtocol(4)
	if ip, err := ip4Consensus.ExternalIP(); err == nil {
		return ip, nil
	}

	// Try IPv6 as fallback
	ip6Consensus := externalip.DefaultConsensus(nil, nil)
	ip6Consensus.UseIPProtocol(6)
	return ip6Consensus.ExternalIP()
}

// Stader client
type Client struct {
	configPath         string
	daemonPath         string
	maxFee             float64
	maxPrioFee         float64
	gasLimit           uint64
	customNonce        *big.Int
	client             *ssh.Client
	originalMaxFee     float64
	originalMaxPrioFee float64
	originalGasLimit   uint64
	debugPrint         bool
	ignoreSyncCheck    bool
	forceFallbacks     bool
}

// Create new Stader client from CLI context
func NewClientFromCtx(c *cli.Context) (*Client, error) {
	return NewClient(c.GlobalString("config-path"),
		c.GlobalString("daemon-path"),
		c.GlobalFloat64("maxFee"),
		c.GlobalFloat64("maxPrioFee"),
		c.GlobalUint64("gasLimit"),
		c.GlobalString("nonce"),
		c.GlobalBool("debug"))
}

// Create new Stader client
func NewClient(configPath string, daemonPath string, maxFee float64, maxPrioFee float64, gasLimit uint64, customNonce string, debug bool) (*Client, error) {

	// Initialize SSH client if configured for SSH
	var sshClient *ssh.Client
	var customNonceBigInt *big.Int = nil
	var success bool
	if customNonce != "" {
		customNonceBigInt, success = big.NewInt(0).SetString(customNonce, 0)
		if !success {
			return nil, fmt.Errorf("Invalid nonce: %s", customNonce)
		}
	}

	// Return client
	client := &Client{
		configPath:         os.ExpandEnv(configPath),
		daemonPath:         os.ExpandEnv(daemonPath),
		maxFee:             maxFee,
		maxPrioFee:         maxPrioFee,
		gasLimit:           gasLimit,
		originalMaxFee:     maxFee,
		originalMaxPrioFee: maxPrioFee,
		originalGasLimit:   gasLimit,
		customNonce:        customNonceBigInt,
		client:             sshClient,
		debugPrint:         debug,
		forceFallbacks:     false,
		ignoreSyncCheck:    false,
	}

	return client, nil

}

// Close client remote connection
func (c *Client) Close() {
	if c.client != nil {
		_ = c.client.Close()
	}
}

// Load the config
func (c *Client) LoadConfig() (*config.StaderConfig, bool, error) {
	settingsFilePath := filepath.Join(c.configPath, SettingsFile)
	expandedPath, err := homedir.Expand(settingsFilePath)
	if err != nil {
		return nil, false, fmt.Errorf("error expanding settings file path: %w", err)
	}

	cfg, err := staderUtils.LoadConfigFromFile(expandedPath)
	if err != nil {
		return nil, false, err
	}

	isNew := false
	if cfg == nil {
		cfg = config.NewStaderConfig(c.configPath)
		isNew = true
	}
	return cfg, isNew, nil
}

// Load the backup config
func (c *Client) LoadBackupConfig() (*config.StaderConfig, error) {
	settingsFilePath := filepath.Join(c.configPath, BackupSettingsFile)
	expandedPath, err := homedir.Expand(settingsFilePath)
	if err != nil {
		return nil, fmt.Errorf("error expanding backup settings file path: %w", err)
	}

	return staderUtils.LoadConfigFromFile(expandedPath)
}

// Save the config
func (c *Client) SaveConfig(cfg *config.StaderConfig) error {
	settingsFilePath := filepath.Join(c.configPath, SettingsFile)
	expandedPath, err := homedir.Expand(settingsFilePath)
	if err != nil {
		return err
	}
	return staderUtils.SaveConfig(cfg, expandedPath)
}

// Remove the upgrade flag file
func (c *Client) RemoveUpgradeFlagFile() error {
	expandedPath, err := homedir.Expand(c.configPath)
	if err != nil {
		return err
	}
	return staderUtils.RemoveUpgradeFlagFile(expandedPath)
}

// Returns whether or not this is the first run of the configurator since a previous installation
func (c *Client) IsFirstRun() (bool, error) {
	expandedPath, err := homedir.Expand(c.configPath)
	if err != nil {
		return false, fmt.Errorf("error expanding settings file path: %w", err)
	}
	return staderUtils.IsFirstRun(expandedPath), nil
}

// Load the Prometheus template, do an environment variable substitution, and save it
func (c *Client) UpdatePrometheusConfiguration(settings map[string]string) error {
	prometheusTemplatePath, err := homedir.Expand(fmt.Sprintf("%s/%s", c.configPath, PrometheusConfigTemplate))
	if err != nil {
		return fmt.Errorf("Error expanding Prometheus template path: %w", err)
	}

	prometheusConfigPath, err := homedir.Expand(fmt.Sprintf("%s/%s", c.configPath, PrometheusFile))
	if err != nil {
		return fmt.Errorf("Error expanding Prometheus config file path: %w", err)
	}

	// Set the environment variables defined in the user settings for metrics
	oldValues := map[string]string{}
	for varName, varValue := range settings {
		oldValues[varName] = os.Getenv(varName)
		os.Setenv(varName, varValue)
	}

	// Read and substitute the template
	contents, err := envsubst.ReadFile(prometheusTemplatePath)
	if err != nil {
		return fmt.Errorf("Error reading and substituting Prometheus configuration template: %w", err)
	}

	// Unset the env vars
	for name, value := range oldValues {
		os.Setenv(name, value)
	}

	// Write the actual Prometheus config file
	err = ioutil.WriteFile(prometheusConfigPath, contents, 0664)
	if err != nil {
		return fmt.Errorf("Could not write Prometheus config file to %s: %w", shellescape.Quote(prometheusConfigPath), err)
	}
	err = os.Chmod(prometheusConfigPath, 0664)
	if err != nil {
		return fmt.Errorf("Could not set Prometheus config file permissions: %w", shellescape.Quote(prometheusConfigPath), err)
	}

	return nil
}

// Install the Stader service
func (c *Client) InstallService(verbose, noDeps bool, network, version, path string, dataPath string) error {

	downloader, err := c.getDownloader()
	if err != nil {
		return err
	}

	// Get installation script flags
	flags := []string{
		"-n", fmt.Sprintf("%s", shellescape.Quote(network)),
		"-v", fmt.Sprintf("%s", shellescape.Quote(version)),
	}
	if path != "" {
		flags = append(flags, fmt.Sprintf("-p %s", shellescape.Quote(path)))
	}
	if noDeps {
		flags = append(flags, "-d")
	}
	if dataPath != "" {
		flags = append(flags, fmt.Sprintf("-u %s", dataPath))
	}

	// Initialize installation command
	cmd, err := c.newCommand(fmt.Sprintf("%s %s | sh -s -- %s", downloader, fmt.Sprintf(InstallerURL, version), strings.Join(flags, " ")))

	if err != nil {
		return err
	}
	defer func() {
		_ = cmd.Close()
	}()

	// fmt.Println("cmd: ", cmd)

	// Get command output pipes
	cmdOut, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmdErr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	// Print progress from stdout
	go (func() {
		scanner := bufio.NewScanner(cmdOut)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	})()

	// Read command & error output from stderr; render in verbose mode
	var errMessage string
	go (func() {
		c := color.New(DebugColor)
		scanner := bufio.NewScanner(cmdErr)
		for scanner.Scan() {
			errMessage = scanner.Text()
			if verbose {
				_, _ = c.Println(scanner.Text())
			}
		}
	})()

	// Run command and return error output
	err = cmd.Run()
	fmt.Println("err: ", err)
	if err != nil {
		return fmt.Errorf("Could not install Stader service: %s", errMessage)
	}
	return nil

}

// Start the Stader service
func (c *Client) StartService(composeFiles []string) error {

	// Start the API container first
	fmt.Println("Starting API container...")
	cmd, err := c.compose([]string{}, "up -d")
	if err != nil {
		return fmt.Errorf("error creating compose command for API container: %w", err)
	}
	err = c.printOutput(cmd)
	if err != nil {
		return fmt.Errorf("error starting API container: %w", err)
	}

	// Start all of the containers
	fmt.Println("Starting Stader containers...")
	cmd, err = c.compose(composeFiles, "up -d --remove-orphans")
	if err != nil {
		return err
	}
	return c.printOutput(cmd)
}

// Pause the Stader service
func (c *Client) PauseService(composeFiles []string) error {
	cmd, err := c.compose(composeFiles, "stop")
	if err != nil {
		return err
	}
	return c.printOutput(cmd)
}

// Stop the Stader service
func (c *Client) StopService(composeFiles []string) error {
	cmd, err := c.compose(composeFiles, "down -v")
	if err != nil {
		return err
	}
	return c.printOutput(cmd)
}

// Stop the Stader service and remove the config folder
func (c *Client) TerminateService(composeFiles []string, configPath string) error {
	// Get the command to run with root privileges
	rootCmd, err := c.getEscalationCommand()
	if err != nil {
		return fmt.Errorf("could not get privilege escalation command: %w", err)
	}

	// Terminate the Docker containers
	cmd, err := c.compose(composeFiles, "down -v")
	if err != nil {
		return fmt.Errorf("error creating Docker artifact removal command: %w", err)
	}
	err = c.printOutput(cmd)
	if err != nil {
		return fmt.Errorf("error removing Docker artifacts: %w", err)
	}

	// Delete the  directory
	path, err := homedir.Expand(configPath)
	if err != nil {
		return fmt.Errorf("error loading Stader directory: %w", err)
	}
	fmt.Printf("Deleting Stader directory (%s)...\n", path)
	cmd = fmt.Sprintf("%s rm -rf %s", rootCmd, path)
	_, err = c.readOutput(cmd)
	if err != nil {
		return fmt.Errorf("error deleting Stader directory: %w", err)
	}

	fmt.Println("Termination complete.")

	return nil
}

// Print the Stader service status
func (c *Client) PrintServiceStatus(composeFiles []string) error {
	cmd, err := c.compose(composeFiles, "ps")
	if err != nil {
		return err
	}
	return c.printOutput(cmd)
}

// Print the Stader service logs
func (c *Client) PrintServiceLogs(composeFiles []string, tail string, serviceNames ...string) error {
	sanitizedStrings := make([]string, len(serviceNames))
	for i, serviceName := range serviceNames {
		sanitizedStrings[i] = fmt.Sprintf("%s", shellescape.Quote(serviceName))
	}
	cmd, err := c.compose(composeFiles, fmt.Sprintf("logs -f --tail %s %s", shellescape.Quote(tail), strings.Join(sanitizedStrings, " ")))
	if err != nil {
		return err
	}
	return c.printOutput(cmd)
}

// Print the Stader service stats
func (c *Client) PrintServiceStats(composeFiles []string) error {

	// Get service container IDs
	cmd, err := c.compose(composeFiles, "ps -q")
	if err != nil {
		return err
	}
	containers, err := c.readOutput(cmd)
	if err != nil {
		return err
	}
	containerIds := strings.Split(strings.TrimSpace(string(containers)), "\n")

	// Print stats
	return c.printOutput(fmt.Sprintf("docker stats %s", strings.Join(containerIds, " ")))

}

// Print the Stader service compose config
func (c *Client) PrintServiceCompose(composeFiles []string) error {
	cmd, err := c.compose(composeFiles, "config")
	if err != nil {
		return err
	}
	return c.printOutput(cmd)
}

// Get the Stader service version
func (c *Client) GetServiceVersion() (string, error) {

	// Get service container version output
	var cmd string
	if c.daemonPath == "" {
		containerName, err := c.getAPIContainerName()
		if err != nil {
			return "", err
		}
		cmd = fmt.Sprintf("docker exec %s %s --version", shellescape.Quote(containerName), shellescape.Quote(APIBinPath))
	} else {
		cmd = fmt.Sprintf("%s --version", shellescape.Quote(c.daemonPath))
	}
	versionBytes, err := c.readOutput(cmd)
	if err != nil {
		return "", fmt.Errorf("Could not get Stader service version: %w", err)
	}

	// Get the version string
	outputString := string(versionBytes)
	elements := strings.Fields(outputString) // Split on whitespace
	if len(elements) < 1 {
		return "", fmt.Errorf("Could not parse Stader service version number from output '%s'", outputString)
	}
	versionString := elements[len(elements)-1]

	// Make sure it's a semantic version
	version, err := semver.Make(versionString)
	if err != nil {
		return "", fmt.Errorf("Could not parse Stader service version number from output '%s': %w", outputString, err)
	}

	// Return the parsed semantic version (extra safety)
	return version.String(), nil

}

// Increments the custom nonce parameter.
// This is used for calls that involve multiple transactions, so they don't all have the same nonce.
func (c *Client) IncrementCustomNonce() {
	c.customNonce.Add(c.customNonce, big.NewInt(1))
}

// Get the current Docker image used by the given container
func (c *Client) GetDockerImage(container string) (string, error) {

	cmd := fmt.Sprintf("docker container inspect --format={{.Config.Image}} %s", container)
	image, err := c.readOutput(cmd)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(image)), nil

}

// Get the current Docker image used by the given container
func (c *Client) GetDockerStatus(container string) (string, error) {

	cmd := fmt.Sprintf("docker container inspect --format={{.State.Status}} %s", container)
	status, err := c.readOutput(cmd)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(status)), nil

}

// Get the time that the given container shut down
func (c *Client) GetDockerContainerShutdownTime(container string) (time.Time, error) {

	cmd := fmt.Sprintf("docker container inspect --format={{.State.FinishedAt}} %s", container)
	finishTimeBytes, err := c.readOutput(cmd)
	if err != nil {
		return time.Time{}, err
	}

	finishTime, err := time.Parse(time.RFC3339, strings.TrimSpace(string(finishTimeBytes)))
	if err != nil {
		return time.Time{}, fmt.Errorf("Error parsing validator container exit time [%s]: %w", string(finishTimeBytes), err)
	}

	return finishTime, nil

}

// Shut down a container
func (c *Client) StopContainer(container string) (string, error) {

	cmd := fmt.Sprintf("docker stop %s", container)
	output, err := c.readOutput(cmd)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil

}

// Start a container
func (c *Client) StartContainer(container string) (string, error) {

	cmd := fmt.Sprintf("docker start %s", container)
	output, err := c.readOutput(cmd)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil

}

// Restart a container
func (c *Client) RestartContainer(container string) (string, error) {

	cmd := fmt.Sprintf("docker restart %s", container)
	output, err := c.readOutput(cmd)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil

}

// Deletes a container
func (c *Client) RemoveContainer(container string) (string, error) {

	cmd := fmt.Sprintf("docker rm %s", container)
	output, err := c.readOutput(cmd)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil

}

// Deletes a container
func (c *Client) DeleteVolume(volume string) (string, error) {

	cmd := fmt.Sprintf("docker volume rm %s", volume)
	output, err := c.readOutput(cmd)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil

}

// Gets the absolute file path of the client volume
func (c *Client) GetClientVolumeSource(container string, volumeTarget string) (string, error) {

	cmd := fmt.Sprintf("docker container inspect --format='{{range .Mounts}}{{if eq \"%s\" .Destination}}{{.Source}}{{end}}{{end}}' %s", volumeTarget, container)
	output, err := c.readOutput(cmd)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// Gets the name of the client volume
func (c *Client) GetClientVolumeName(container string, volumeTarget string) (string, error) {

	cmd := fmt.Sprintf("docker container inspect --format='{{range .Mounts}}{{if eq \"%s\" .Destination}}{{.Name}}{{end}}{{end}}' %s", volumeTarget, container)
	output, err := c.readOutput(cmd)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// Gets the disk usage of the given volume
func (c *Client) GetVolumeSize(volumeName string) (string, error) {

	cmd := fmt.Sprintf("docker system df -v --format='{{range .Volumes}}{{if eq \"%s\" .Name}}{{.Size}}{{end}}{{end}}'", volumeName)
	output, err := c.readOutput(cmd)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

// Runs the prune provisioner
func (c *Client) RunPruneProvisioner(container string, volume string, image string) error {

	// Run the prune provisioner
	cmd := fmt.Sprintf("docker run --rm --name %s -v %s:/ethclient %s", container, volume, image)
	output, err := c.readOutput(cmd)
	if err != nil {
		return err
	}

	outputString := strings.TrimSpace(string(output))
	if outputString != "" {
		return fmt.Errorf("Unexpected output running the prune provisioner: %s", outputString)
	}

	return nil

}

// Runs the prune provisioner
func (c *Client) RunNethermindPruneStarter(container string) error {
	cmd := fmt.Sprintf("docker exec %s %s %s", container, nethermindPruneStarterCommand, nethermindAdminUrl)
	err := c.printOutput(cmd)
	if err != nil {
		return err
	}
	return nil
}

// Runs the EC migrator
func (c *Client) RunEcMigrator(container string, volume string, targetDir string, mode string, image string) error {
	cmd := fmt.Sprintf("docker run --rm --name %s -v %s:/ethclient -v %s:/mnt/external -e EC_MIGRATE_MODE='%s' %s", container, volume, targetDir, mode, image)
	err := c.printOutput(cmd)
	if err != nil {
		return err
	}

	return nil
}

// Gets the size of the target directory via the EC migrator for importing, which should have the same permissions as exporting
func (c *Client) GetDirSizeViaEcMigrator(container string, targetDir string, image string) (uint64, error) {
	cmd := fmt.Sprintf("docker run --rm --name %s -v %s:/mnt/external -e OPERATION='size' %s", container, targetDir, image)
	output, err := c.readOutput(cmd)
	if err != nil {
		return 0, fmt.Errorf("Error getting source directory size: %w", err)
	}

	trimmedOutput := strings.TrimRight(string(output), "\n")
	dirSize, err := strconv.ParseUint(trimmedOutput, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("Error parsing directory size output [%s]: %w", trimmedOutput, err)
	}

	return dirSize, nil
}

// Deletes the node wallet and all validator keys, and restarts the Docker containers
func (c *Client) PurgeAllKeys(composeFiles []string) error {
	// Get the command to run with root privileges
	rootCmd, err := c.getEscalationCommand()
	if err != nil {
		return fmt.Errorf("could not get privilege escalation command: %w", err)
	}

	// Get the config
	cfg, _, err := c.LoadConfig()
	if err != nil {
		return fmt.Errorf("error loading user settings: %w", err)
	}

	// Shut down the containers
	fmt.Println("Stopping containers...")
	err = c.PauseService(composeFiles)
	if err != nil {
		return fmt.Errorf("error stopping Docker containers: %w", err)
	}

	// Delete the wallet
	walletPath, err := homedir.Expand(cfg.StaderNode.GetWalletPathInCLI())
	if err != nil {
		return fmt.Errorf("error loading wallet path: %w", err)
	}
	fmt.Println("Deleting wallet...")
	cmd := fmt.Sprintf("%s rm -f %s", rootCmd, walletPath)
	_, err = c.readOutput(cmd)
	if err != nil {
		return fmt.Errorf("error deleting wallet: %w", err)
	}

	// Delete the password
	passwordPath, err := homedir.Expand(cfg.StaderNode.GetPasswordPathInCLI())
	if err != nil {
		return fmt.Errorf("error loading password path: %w", err)
	}
	fmt.Println("Deleting password...")
	cmd = fmt.Sprintf("%s rm -f %s", rootCmd, passwordPath)
	_, err = c.readOutput(cmd)
	if err != nil {
		return fmt.Errorf("error deleting password: %w", err)
	}

	// Delete the validators dir
	validatorsPath, err := homedir.Expand(cfg.StaderNode.GetValidatorKeychainPathInCLI())
	if err != nil {
		return fmt.Errorf("error loading validators folder path: %w", err)
	}
	fmt.Println("Deleting validator keys...")
	cmd = fmt.Sprintf("%s rm -rf %s/*", rootCmd, validatorsPath)
	_, err = c.readOutput(cmd)
	if err != nil {
		return fmt.Errorf("error deleting validator keys: %w", err)
	}
	cmd = fmt.Sprintf("%s rm -rf %s/.[a-zA-Z0-9]*", rootCmd, validatorsPath)
	_, err = c.readOutput(cmd)
	if err != nil {
		return fmt.Errorf("error deleting hidden files in validator folder: %w", err)
	}

	// Start the containers
	fmt.Println("Starting containers...")
	err = c.StartService(composeFiles)
	if err != nil {
		return fmt.Errorf("error starting Docker containers: %w", err)
	}

	fmt.Println("Purge complete.")

	return nil
}

// Get the gas settings
func (c *Client) GetGasSettings() (float64, float64, uint64) {
	return c.maxFee, c.maxPrioFee, c.gasLimit
}

// Get the gas fees
func (c *Client) AssignGasSettings(maxFee float64, maxPrioFee float64, gasLimit uint64) {
	c.maxFee = maxFee
	c.maxPrioFee = maxPrioFee
	c.gasLimit = gasLimit
}

// Set the flags for ignoring EC and CC sync checks and forcing fallbacks to prevent unnecessary duplication of effort by the API during CLI commands
func (c *Client) SetClientStatusFlags(ignoreSyncCheck bool, forceFallbacks bool) {
	c.ignoreSyncCheck = ignoreSyncCheck
	c.forceFallbacks = forceFallbacks
}

// Get the command used to escalate privileges on the system
func (c *Client) getEscalationCommand() (string, error) {
	// Check for sudo first
	sudo := "sudo"
	exists, err := c.checkIfCommandExists(sudo)
	if err != nil {
		return "", fmt.Errorf("error checking if %s exists: %w", sudo, err)
	}
	if exists {
		return sudo, nil
	}

	// Check for doas next
	doas := "doas"
	exists, err = c.checkIfCommandExists(doas)
	if err != nil {
		return "", fmt.Errorf("error checking if %s exists: %w", doas, err)
	}
	if exists {
		return doas, nil
	}

	return "", fmt.Errorf("no privilege escalation command found")
}

func (c *Client) checkIfCommandExists(command string) (bool, error) {
	// Run `type` to check for existence
	cmd := fmt.Sprintf("type %s", command)
	output, err := c.readOutput(cmd)

	if err != nil {
		exitErr, isExitErr := err.(*exec.ExitError)
		if isExitErr && exitErr.ProcessState.ExitCode() == 127 {
			// Command not found
			return false, nil
		} else {
			return false, fmt.Errorf("error checking if %s exists: %w", command, err)
		}
	} else {
		if strings.Contains(string(output), fmt.Sprintf("%s is", command)) {
			return true, nil
		} else {
			return false, fmt.Errorf("unexpected output when checking for %s: %s", command, string(output))
		}
	}
}

// Get the provider mode and port from a legacy config's provider URL
func (c *Client) migrateProviderInfo(provider string, wsProvider string, localHostname string, clientMode *cfgtypes.Parameter, httpPortParam *cfgtypes.Parameter, wsPortParam *cfgtypes.Parameter, externalHttpUrlParam *cfgtypes.Parameter, externalWsUrlParam *cfgtypes.Parameter) error {

	// Get HTTP provider
	mode, port, err := c.getLegacyProviderInfo(provider, localHostname)
	if err != nil {
		return fmt.Errorf("error parsing %s provider: %w", localHostname, err)
	}

	// Set the mode, provider, port, and/or URL
	clientMode.Value = mode
	if mode == cfgtypes.Mode_Local {
		httpPortParam.Value = port
	} else {
		externalHttpUrlParam.Value = provider
	}

	// Get the websocket provider
	if wsProvider != "" {
		_, wsPort, err := c.getLegacyProviderInfo(wsProvider, localHostname)
		if err != nil {
			return fmt.Errorf("error parsing %s websocket provider: %w", localHostname, err)
		}
		if mode == cfgtypes.Mode_Local {
			wsPortParam.Value = wsPort
		} else {
			externalWsUrlParam.Value = wsProvider
		}
	}

	return nil

}

// Get the provider mode and port from a legacy config's provider URL
func (c *Client) getLegacyProviderInfo(provider string, localHostname string) (cfgtypes.Mode, uint16, error) {

	providerUrl, err := url.Parse(provider)
	if err != nil {
		return cfgtypes.Mode_Unknown, 0, fmt.Errorf("error parsing %s provider: %w", localHostname, err)
	}

	var mode cfgtypes.Mode
	if providerUrl.Hostname() == localHostname {
		// This is Docker mode
		mode = cfgtypes.Mode_Local
	} else {
		// This is Hybrid mode
		mode = cfgtypes.Mode_External
	}

	var port uint16
	portString := providerUrl.Port()
	if portString == "" {
		switch providerUrl.Scheme {
		case "http", "ws":
			port = 80
		case "https", "wss":
			port = 443
		default:
			return cfgtypes.Mode_Unknown, 0, fmt.Errorf("provider [%s] doesn't provide port info and it can't be inferred from the scheme", provider)
		}
	} else {
		parsedPort, err := strconv.ParseUint(portString, 0, 16)
		if err != nil {
			return cfgtypes.Mode_Unknown, 0, fmt.Errorf("invalid port [%s] in %s provider [%s]", portString, localHostname, provider)
		}
		port = uint16(parsedPort)
	}

	return mode, port, nil

}

// Sets a modern config's selected EC / mode based on a legacy config
func (c *Client) migrateEcSelection(legacySelectedClient string, ecParam *cfgtypes.Parameter, ecModeParam *cfgtypes.Parameter) error {
	// EC selection
	switch legacySelectedClient {
	case "geth":
		ecParam.Value = cfgtypes.ExecutionClient_Geth
	case "infura":
		ecParam.Value = cfgtypes.ExecutionClient_Geth
	case "pocket":
		ecParam.Value = cfgtypes.ExecutionClient_Geth
	case "custom":
		ecModeParam.Value = cfgtypes.Mode_External
	case "":
		break
	default:
		return fmt.Errorf("unknown eth1 client [%s]", legacySelectedClient)
	}

	return nil
}

// Sets a modern config's selected CC / mode based on a legacy config
func (c *Client) migrateCcSelection(legacySelectedClient string, ccParam *cfgtypes.Parameter) error {
	// CC selection
	switch legacySelectedClient {
	case "lighthouse":
		ccParam.Value = cfgtypes.ConsensusClient_Lighthouse
	case "nimbus":
		ccParam.Value = cfgtypes.ConsensusClient_Nimbus
	case "prysm":
		ccParam.Value = cfgtypes.ConsensusClient_Prysm
	case "teku":
		ccParam.Value = cfgtypes.ConsensusClient_Teku
	default:
		return fmt.Errorf("unknown eth2 client [%s]", legacySelectedClient)
	}

	return nil
}

// Build a docker compose command
func (c *Client) compose(composeFiles []string, args string) (string, error) {

	// Cancel if running in non-docker mode
	if c.daemonPath != "" {
		return "", errors.New("command unavailable in Native Mode (with '--daemon-path' option specified)")
	}

	// Get the expanded config path
	expandedConfigPath, err := homedir.Expand(c.configPath)
	if err != nil {
		return "", err
	}

	// Load config
	cfg, isNew, err := c.LoadConfig()
	if err != nil {
		return "", err
	}

	if isNew {
		return "", fmt.Errorf("Settings file not found. Please run `stader-permissioned-cli service config` to set up your Stadernode before starting it.")
	}

	// Get the external IP address
	var externalIP string
	ip, err := getExternalIP()
	if err != nil {
		fmt.Println("Warning: couldn't get external IP address; if you're using Nimbus or Besu, it may have trouble finding peers:")
		fmt.Println(err.Error())
	} else {
		if ip.To4() == nil {
			fmt.Println("Warning: external IP address is v6; if you're using Nimbus or Besu, it may have trouble finding peers:")
		}
		externalIP = ip.String()
	}

	// Set up environment variables and deploy the template config files
	settings := cfg.GenerateEnvironmentVariables()
	settings["EXTERNAL_IP"] = shellescape.Quote(externalIP)

	// Deploy the templates and run environment variable substitution on them
	deployedContainers, err := c.deployTemplates(cfg, expandedConfigPath, settings)
	if err != nil {
		return "", fmt.Errorf("error deploying Docker templates: %w", err)
	}

	// Set up all of the environment variables to pass to the run command
	env := []string{}
	for key, value := range settings {
		env = append(env, fmt.Sprintf("%s=%s", key, shellescape.Quote(value)))
	}

	// Include all of the relevant docker compose definition files
	composeFileFlags := []string{}
	for _, container := range deployedContainers {
		composeFileFlags = append(composeFileFlags, fmt.Sprintf("-f %s", shellescape.Quote(container)))
	}
	for _, container := range composeFiles {
		composeFileFlags = append(composeFileFlags, fmt.Sprintf("-f %s", shellescape.Quote(container)))
	}

	// Return command
	return fmt.Sprintf("%s docker compose --project-directory %s %s %s", strings.Join(env, " "), shellescape.Quote(expandedConfigPath), strings.Join(composeFileFlags, " "), args), nil

}

// Deploys all of the appropriate docker compose template files and provisions them based on the provided configuration
func (c *Client) deployTemplates(cfg *config.StaderConfig, staderDir string, settings map[string]string) ([]string, error) {

	// Check for the folders
	runtimeFolder := filepath.Join(staderDir, runtimeDir)
	templatesFolder := filepath.Join(staderDir, templatesDir)
	_, err := os.Stat(templatesFolder)
	if os.IsNotExist(err) {
		return []string{}, fmt.Errorf("templates folder [%s] does not exist", templatesFolder)
	}
	overrideFolder := filepath.Join(staderDir, overrideDir)
	_, err = os.Stat(overrideFolder)
	if os.IsNotExist(err) {
		return []string{}, fmt.Errorf("override folder [%s] does not exist", overrideFolder)
	}

	// Clear out the runtime folder and remake it
	err = os.RemoveAll(runtimeFolder)
	if err != nil {
		return []string{}, fmt.Errorf("error deleting runtime folder [%s]: %w", runtimeFolder, err)
	}
	err = os.Mkdir(runtimeFolder, 0775)
	if err != nil {
		return []string{}, fmt.Errorf("error creating runtime folder [%s]: %w", runtimeFolder, err)
	}

	// Set the environment variables for substitution
	oldValues := map[string]string{}
	for varName, varValue := range settings {
		oldValues[varName] = os.Getenv(varName)
		os.Setenv(varName, varValue)
	}
	defer func() {
		// Unset the env vars
		for name, value := range oldValues {
			os.Setenv(name, value)
		}
	}()

	// Read and substitute the templates
	deployedContainers := []string{}

	// return deployedContainers, nil
	// API
	contents, err := envsubst.ReadFile(filepath.Join(templatesFolder, config.ApiContainerName+templateSuffix))
	if err != nil {
		return []string{}, fmt.Errorf("error reading and substituting API container template: %w", err)
	}
	apiComposePath := filepath.Join(runtimeFolder, config.ApiContainerName+composeFileSuffix)
	err = ioutil.WriteFile(apiComposePath, contents, 0664)
	if err != nil {
		return []string{}, fmt.Errorf("could not write API container file to %s: %w", apiComposePath, err)
	}
	deployedContainers = append(deployedContainers, apiComposePath)
	deployedContainers = append(deployedContainers, filepath.Join(overrideFolder, config.ApiContainerName+composeFileSuffix))

	// Node
	contents, err = envsubst.ReadFile(filepath.Join(templatesFolder, config.NodeContainerName+templateSuffix))
	if err != nil {
		return []string{}, fmt.Errorf("error reading and substituting node container template: %w", err)
	}
	nodeComposePath := filepath.Join(runtimeFolder, config.NodeContainerName+composeFileSuffix)
	err = ioutil.WriteFile(nodeComposePath, contents, 0664)
	if err != nil {
		return []string{}, fmt.Errorf("could not write node container file to %s: %w", nodeComposePath, err)
	}
	deployedContainers = append(deployedContainers, nodeComposePath)
	deployedContainers = append(deployedContainers, filepath.Join(overrideFolder, config.NodeContainerName+composeFileSuffix))

	return c.composeAddons(cfg, staderDir, settings, deployedContainers)

}

// Handle composing for addons
func (c *Client) composeAddons(cfg *config.StaderConfig, staderDir string, settings map[string]string, deployedContainers []string) ([]string, error) {

	return deployedContainers, nil

}

// Call the Stader API
func (c *Client) callAPI(args string, otherArgs ...string) ([]byte, error) {
	// Sanitize and parse the args
	ignoreSyncCheckFlag, forceFallbackECFlag, args := c.getApiCallArgs(args, otherArgs...)

	// Create the command to run
	var cmd string
	if c.daemonPath == "" {
		containerName, err := c.getAPIContainerName()
		if err != nil {
			return []byte{}, err
		}
		cmd = fmt.Sprintf("docker exec %s %s %s %s %s %s api %s", shellescape.Quote(containerName), shellescape.Quote(APIBinPath), ignoreSyncCheckFlag, forceFallbackECFlag, c.getGasOpts(), c.getCustomNonce(), args)
	} else {
		cmd = fmt.Sprintf("%s --settings %s %s %s %s %s api %s",
			c.daemonPath,
			shellescape.Quote(fmt.Sprintf("%s/%s", c.configPath, SettingsFile)),
			ignoreSyncCheckFlag,
			forceFallbackECFlag,
			c.getGasOpts(),
			c.getCustomNonce(),
			args)
	}

	// Run the command
	return c.runApiCall(cmd)
}

// Call the Stader API with some custom environment variables
func (c *Client) callAPIWithEnvVars(envVars map[string]string, args string, otherArgs ...string) ([]byte, error) {
	// Sanitize and parse the args
	ignoreSyncCheckFlag, forceFallbackECFlag, args := c.getApiCallArgs(args, otherArgs...)

	// Create the command to run
	var cmd string
	if c.daemonPath == "" {
		envArgs := ""
		for key, value := range envVars {
			os.Setenv(key, shellescape.Quote(value))
			envArgs += fmt.Sprintf("-e %s ", key)
		}
		containerName, err := c.getAPIContainerName()
		if err != nil {
			return []byte{}, err
		}
		cmd = fmt.Sprintf("docker exec %s %s %s %s %s %s %s api %s", envArgs, shellescape.Quote(containerName), shellescape.Quote(APIBinPath), ignoreSyncCheckFlag, forceFallbackECFlag, c.getGasOpts(), c.getCustomNonce(), args)
	} else {
		envArgs := ""
		for key, value := range envVars {
			envArgs += fmt.Sprintf("%s=%s ", key, shellescape.Quote(value))
		}
		cmd = fmt.Sprintf("%s %s --settings %s %s %s %s %s api %s",
			envArgs,
			c.daemonPath,
			shellescape.Quote(fmt.Sprintf("%s/%s", c.configPath, SettingsFile)),
			ignoreSyncCheckFlag,
			forceFallbackECFlag,
			c.getGasOpts(),
			c.getCustomNonce(),
			args)
	}

	// Run the command
	return c.runApiCall(cmd)
}

func (c *Client) getApiCallArgs(args string, otherArgs ...string) (string, string, string) {
	// Sanitize arguments
	var sanitizedArgs []string
	for _, arg := range strings.Fields(args) {
		sanitizedArg := shellescape.Quote(arg)
		sanitizedArgs = append(sanitizedArgs, sanitizedArg)
	}
	args = strings.Join(sanitizedArgs, " ")
	if len(otherArgs) > 0 {
		for _, arg := range otherArgs {
			sanitizedArg := shellescape.Quote(arg)
			args += fmt.Sprintf(" %s", sanitizedArg)
		}
	}

	ignoreSyncCheckFlag := ""
	if c.ignoreSyncCheck {
		ignoreSyncCheckFlag = "--ignore-sync-check"
	}
	forceFallbacksFlag := ""
	if c.forceFallbacks {
		forceFallbacksFlag = "--force-fallbacks"
	}

	return ignoreSyncCheckFlag, forceFallbacksFlag, args
}

func (c *Client) runApiCall(cmd string) ([]byte, error) {
	if c.debugPrint {
		fmt.Println("To API:")
		fmt.Println(cmd)
	}

	output, err := c.readOutput(cmd)

	if c.debugPrint {
		if output != nil {
			fmt.Println("API Out:")
			fmt.Println(string(output))
		}
		if err != nil {
			fmt.Println("API Err:")
			fmt.Println(err.Error())
		}
	}

	// Reset the gas settings after the call
	c.maxFee = c.originalMaxFee
	c.maxPrioFee = c.originalMaxPrioFee
	c.gasLimit = c.originalGasLimit

	return output, err
}

// Get the API container name
func (c *Client) getAPIContainerName() (string, error) {
	cfg, _, err := c.LoadConfig()
	if err != nil {
		return "", err
	}
	if cfg.StaderNode.ProjectName.Value == "" {
		return "", errors.New("Stader docker project name not set")
	}
	return cfg.StaderNode.ProjectName.Value.(string) + APIContainerSuffix, nil
}

// Get gas price & limit flags
func (c *Client) getGasOpts() string {
	var opts string
	opts += fmt.Sprintf("--maxFee %f ", c.maxFee)
	opts += fmt.Sprintf("--maxPrioFee %f ", c.maxPrioFee)
	opts += fmt.Sprintf("--gasLimit %d ", 100000)
	return opts
}

func (c *Client) getCustomNonce() string {
	// Set the custom nonce
	nonce := ""
	if c.customNonce != nil {
		nonce = fmt.Sprintf("--nonce %s", c.customNonce.String())
	}
	return nonce
}

// Get the first downloader available to the system
func (c *Client) getDownloader() (string, error) {

	// Check for cURL
	hasCurl, err := c.readOutput("command -v curl")
	if err == nil && len(hasCurl) > 0 {
		return "curl -sL", nil
	}

	// Check for wget
	hasWget, err := c.readOutput("command -v wget")
	if err == nil && len(hasWget) > 0 {
		return "wget -qO-", nil
	}

	// Return error
	return "", errors.New("Either cURL or wget is required to begin installation.")

}

// pipeToStdOut pipes cmdOut to stdout
func pipeToStdOut(cmdOut io.Reader) {
	_, err := io.Copy(os.Stdout, cmdOut)
	if err != nil {
		log.Printf("Error piping stdout: %v", err)
	}
}

// pipeToStdErr pipes cmdErr to stderr
func pipeToStdErr(cmdErr io.Reader) {
	_, err := io.Copy(os.Stderr, cmdErr)
	if err != nil {
		log.Printf("Error piping stderr: %v", err)
	}
}

// pipeOutput pipes cmdOut and cmdErr to stdout and stderr
func pipeOutput(cmdOut, cmdErr io.Reader) {
	go pipeToStdOut(cmdOut)
	go pipeToStdErr(cmdErr)
}

// Run a command and print its output
func (c *Client) printOutput(cmdText string) error {

	// Initialize command
	cmd, err := c.newCommand(cmdText)
	if err != nil {
		return err
	}
	defer cmd.Close()

	cmdOut, cmdErr, err := cmd.OutputPipes()
	if err != nil {
		return err
	}

	// Begin piping before the command is started
	pipeOutput(cmdOut, cmdErr)

	// Start the command
	if err := cmd.Start(); err != nil {
		return err
	}

	// Wait for the command to exit
	return cmd.Wait()

}

// Run a command and return its output
func (c *Client) readOutput(cmdText string) ([]byte, error) {

	// Initialize command
	cmd, err := c.newCommand(cmdText)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		_ = cmd.Close()
	}()

	// Run command and return output
	return cmd.Output()

}
