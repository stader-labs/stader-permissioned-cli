/*
This work is licensed and released under GNU GPL v3 or any other later versions.
The full text of the license is below/ found at <http://www.gnu.org/licenses/>

(c) 2023 Rocket Pool Pty Ltd. Modified under GNU GPL v3. [1.0.0]

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
package wallet

import (
	"fmt"

	"github.com/stader-labs/stader-node/shared/services"
	"github.com/stader-labs/stader-node/shared/types/api"
	"github.com/urfave/cli"
)

func purge(c *cli.Context) (*api.PurgeResponse, error) {

	w, err := services.GetWallet(c)
	if err != nil {
		return nil, err
	}

	pm, err := services.GetPasswordManager(c)
	if err != nil {
		return nil, err
	}

	response := api.PurgeResponse{}

	// Delete the VC directories
	err = w.DeleteValidatorStores()
	if err != nil {
		return nil, fmt.Errorf("error deleting validator storage: %w", err)
	}

	// Delete the wallet and password
	err = w.Delete()
	if err != nil {
		return nil, fmt.Errorf("error deleting wallet: %w", err)
	}
	err = pm.DeletePassword()
	if err != nil {
		return nil, fmt.Errorf("error deleting password: %w", err)
	}

	return &response, nil
}
