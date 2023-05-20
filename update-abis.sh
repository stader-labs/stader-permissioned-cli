# enter in stader-lib/contracts and throw and error if it fails
cd ./stader-lib/contracts;

abigen --abi ./../../abis/StaderVaultFactory.abi.json --pkg contracts --type VaultFactory --out vault-factory.go;
abigen --abi ./../../abis/ValidatorWithdrawVault.abi.json --pkg contracts --type ValidatorWithdrawVault --out validator-withdraw-vault.go;
abigen --abi ./../../abis/StakePoolManager.abi.json --pkg contracts --type StakePoolManager --out stake-pool-manager.go;
abigen --abi ./../../abis/StaderConfig.abi.json --pkg contracts --type StaderConfig --out stader-config.go;
abigen --abi ./../../abis/SocializingPool.json --pkg contracts --type SocializingPool --out socializing-pool.go;
abigen --abi ./../../abis/PoolUtils.abi.json --pkg contracts --type PoolUtils --out pool-utils.go;
abigen --abi ./../../abis/PermissionedPool.abi.json --pkg contracts --type PermissionedPool --out permissioned-pool.go;
abigen --abi ./../../abis/PermissionedNodeRegistry.abi.json --pkg contracts --type PermissionedNodeRegistry --out permissioned-node-registry.go;
abigen --abi ./../../abis/Penalty.abi.json --pkg contracts --type PenaltyTracker --out penalty.go;
abigen --abi ./../../abis/NodeElRewardVault.abi.json --pkg contracts --type NodeElRewardVault --out node-el-reward-vault.go;