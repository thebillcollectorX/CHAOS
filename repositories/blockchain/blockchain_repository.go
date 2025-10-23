package blockchain

import (
	"github.com/tiagorlampert/CHAOS/entities"
)

type Repository interface {
	// Network operations
	CreateNetwork(network *entities.Network) error
	GetNetworkByID(id uint) (*entities.Network, error)
	GetNetworkByChainID(chainID uint64) (*entities.Network, error)
	GetActiveNetworks() ([]entities.Network, error)
	UpdateNetwork(network *entities.Network) error
	
	// Wallet operations
	CreateWallet(wallet *entities.Wallet) error
	GetWalletsByUserID(userID uint) ([]entities.Wallet, error)
	GetWalletByAddress(address string) (*entities.Wallet, error)
	UpdateWallet(wallet *entities.Wallet) error
	DeleteWallet(id uint) error
	
	// Transaction operations
	CreateTransaction(tx *entities.Transaction) error
	GetTransactionByHash(hash string) (*entities.Transaction, error)
	GetTransactionsByUserID(userID uint) ([]entities.Transaction, error)
	GetTransactionsByTokenID(tokenID uint) ([]entities.Transaction, error)
	UpdateTransaction(tx *entities.Transaction) error
	
	// Contract Template operations
	CreateTemplate(template *entities.ContractTemplate) error
	GetTemplateByID(id uint) (*entities.ContractTemplate, error)
	GetActiveTemplates() ([]entities.ContractTemplate, error)
	UpdateTemplate(template *entities.ContractTemplate) error
}