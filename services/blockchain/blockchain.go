package blockchain

import "math/big"

type BlockchainService interface {
	DeployContract(network, privateKey, contractCode string, constructorArgs []interface{}) (string, string, error)
	GetGasPrice(network string) (*big.Int, error)
	GetBalance(network, address string) (*big.Int, error)
	WaitForTransaction(network, txHash string) (*types.Receipt, error)
	IsValidAddress(address string) bool
	GetNetworkInfo(network string) (NetworkInfo, error)
}

type NetworkInfo struct {
	Name           string
	Symbol         string
	ChainID        int64
	RPCURL         string
	ExplorerURL    string
	GasPrice       *big.Int
	DeploymentCost *big.Int
}