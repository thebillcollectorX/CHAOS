package blockchain

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

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

type blockchainService struct {
	networks map[string]NetworkInfo
}

func NewBlockchainService() BlockchainService {
	networks := map[string]NetworkInfo{
		"ethereum": {
			Name:           "Ethereum Mainnet",
			Symbol:         "ETH",
			ChainID:        1,
			RPCURL:         "https://mainnet.infura.io/v3/YOUR_PROJECT_ID", // Replace with actual RPC URL
			ExplorerURL:    "https://etherscan.io",
			GasPrice:       big.NewInt(20000000000), // 20 Gwei
			DeploymentCost: big.NewInt(50000000000000000), // 0.05 ETH
		},
		"bsc": {
			Name:           "Binance Smart Chain",
			Symbol:         "BNB",
			ChainID:        56,
			RPCURL:         "https://bsc-dataseed.binance.org/",
			ExplorerURL:    "https://bscscan.com",
			GasPrice:       big.NewInt(5000000000), // 5 Gwei
			DeploymentCost: big.NewInt(10000000000000000), // 0.01 BNB
		},
		"polygon": {
			Name:           "Polygon",
			Symbol:         "MATIC",
			ChainID:        137,
			RPCURL:         "https://polygon-rpc.com/",
			ExplorerURL:    "https://polygonscan.com",
			GasPrice:       big.NewInt(30000000000), // 30 Gwei
			DeploymentCost: big.NewInt(10000000000000000), // 0.01 MATIC
		},
		"arbitrum": {
			Name:           "Arbitrum One",
			Symbol:         "ETH",
			ChainID:        42161,
			RPCURL:         "https://arb1.arbitrum.io/rpc",
			ExplorerURL:    "https://arbiscan.io",
			GasPrice:       big.NewInt(1000000000), // 1 Gwei
			DeploymentCost: big.NewInt(10000000000000000), // 0.01 ETH
		},
		"optimism": {
			Name:           "Optimism",
			Symbol:         "ETH",
			ChainID:        10,
			RPCURL:         "https://mainnet.optimism.io",
			ExplorerURL:    "https://optimistic.etherscan.io",
			GasPrice:       big.NewInt(1000000000), // 1 Gwei
			DeploymentCost: big.NewInt(10000000000000000), // 0.01 ETH
		},
	}

	return &blockchainService{
		networks: networks,
	}
}

func (s *blockchainService) DeployContract(network, privateKey, contractCode string, constructorArgs []interface{}) (string, string, error) {
	networkInfo, exists := s.networks[network]
	if !exists {
		return "", "", fmt.Errorf("unsupported network: %s", network)
	}

	// Connect to the network
	client, err := ethclient.Dial(networkInfo.RPCURL)
	if err != nil {
		return "", "", fmt.Errorf("failed to connect to %s: %v", network, err)
	}
	defer client.Close()

	// Parse private key
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", "", fmt.Errorf("invalid private key: %v", err)
	}

	// Get public key and address
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Get nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "", "", fmt.Errorf("failed to get nonce: %v", err)
	}

	// Get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		gasPrice = networkInfo.GasPrice
	}

	// Create transaction
	tx := types.NewContractCreation(nonce, big.NewInt(0), 500000, gasPrice, []byte(contractCode))

	// Sign transaction
	chainID := big.NewInt(networkInfo.ChainID)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKeyECDSA)
	if err != nil {
		return "", "", fmt.Errorf("failed to sign transaction: %v", err)
	}

	// Send transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", "", fmt.Errorf("failed to send transaction: %v", err)
	}

	return signedTx.Hash().Hex(), networkInfo.ExplorerURL, nil
}

func (s *blockchainService) GetGasPrice(network string) (*big.Int, error) {
	networkInfo, exists := s.networks[network]
	if !exists {
		return nil, fmt.Errorf("unsupported network: %s", network)
	}

	client, err := ethclient.Dial(networkInfo.RPCURL)
	if err != nil {
		return networkInfo.GasPrice, nil // Return default if connection fails
	}
	defer client.Close()

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return networkInfo.GasPrice, nil // Return default if suggestion fails
	}

	return gasPrice, nil
}

func (s *blockchainService) GetBalance(network, address string) (*big.Int, error) {
	networkInfo, exists := s.networks[network]
	if !exists {
		return nil, fmt.Errorf("unsupported network: %s", network)
	}

	client, err := ethclient.Dial(networkInfo.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %v", network, err)
	}
	defer client.Close()

	account := common.HexToAddress(address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %v", err)
	}

	return balance, nil
}

func (s *blockchainService) WaitForTransaction(network, txHash string) (*types.Receipt, error) {
	networkInfo, exists := s.networks[network]
	if !exists {
		return nil, fmt.Errorf("unsupported network: %s", network)
	}

	client, err := ethclient.Dial(networkInfo.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %v", network, err)
	}
	defer client.Close()

	// Wait for transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), client, common.HexToHash(txHash))
	if err != nil {
		return nil, fmt.Errorf("failed to wait for transaction: %v", err)
	}

	return receipt, nil
}

func (s *blockchainService) IsValidAddress(address string) bool {
	return common.IsHexAddress(address)
}

func (s *blockchainService) GetNetworkInfo(network string) (NetworkInfo, error) {
	networkInfo, exists := s.networks[network]
	if !exists {
		return NetworkInfo{}, fmt.Errorf("unsupported network: %s", network)
	}
	return networkInfo, nil
}

// Helper function to convert hex string to bytes
func hexToBytes(hexStr string) ([]byte, error) {
	if strings.HasPrefix(hexStr, "0x") {
		hexStr = hexStr[2:]
	}
	return hex.DecodeString(hexStr)
}