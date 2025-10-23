package blockchain

import (
	"github.com/tiagorlampert/CHAOS/entities"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Network operations
func (r *repository) CreateNetwork(network *entities.Network) error {
	return r.db.Create(network).Error
}

func (r *repository) GetNetworkByID(id uint) (*entities.Network, error) {
	var network entities.Network
	err := r.db.First(&network, id).Error
	if err != nil {
		return nil, err
	}
	return &network, nil
}

func (r *repository) GetNetworkByChainID(chainID uint64) (*entities.Network, error) {
	var network entities.Network
	err := r.db.Where("chain_id = ?", chainID).First(&network).Error
	if err != nil {
		return nil, err
	}
	return &network, nil
}

func (r *repository) GetActiveNetworks() ([]entities.Network, error) {
	var networks []entities.Network
	err := r.db.Where("is_active = ?", true).Find(&networks).Error
	return networks, err
}

func (r *repository) UpdateNetwork(network *entities.Network) error {
	return r.db.Save(network).Error
}

// Wallet operations
func (r *repository) CreateWallet(wallet *entities.Wallet) error {
	return r.db.Create(wallet).Error
}

func (r *repository) GetWalletsByUserID(userID uint) ([]entities.Wallet, error) {
	var wallets []entities.Wallet
	err := r.db.Where("user_id = ?", userID).Find(&wallets).Error
	return wallets, err
}

func (r *repository) GetWalletByAddress(address string) (*entities.Wallet, error) {
	var wallet entities.Wallet
	err := r.db.Where("address = ?", address).First(&wallet).Error
	if err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *repository) UpdateWallet(wallet *entities.Wallet) error {
	return r.db.Save(wallet).Error
}

func (r *repository) DeleteWallet(id uint) error {
	return r.db.Delete(&entities.Wallet{}, id).Error
}

// Transaction operations
func (r *repository) CreateTransaction(tx *entities.Transaction) error {
	return r.db.Create(tx).Error
}

func (r *repository) GetTransactionByHash(hash string) (*entities.Transaction, error) {
	var tx entities.Transaction
	err := r.db.Where("hash = ?", hash).Preload("Token").Preload("User").Preload("Network").First(&tx).Error
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

func (r *repository) GetTransactionsByUserID(userID uint) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := r.db.Where("user_id = ?", userID).Preload("Token").Preload("Network").Find(&transactions).Error
	return transactions, err
}

func (r *repository) GetTransactionsByTokenID(tokenID uint) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := r.db.Where("token_id = ?", tokenID).Preload("User").Preload("Network").Find(&transactions).Error
	return transactions, err
}

func (r *repository) UpdateTransaction(tx *entities.Transaction) error {
	return r.db.Save(tx).Error
}

// Contract Template operations
func (r *repository) CreateTemplate(template *entities.ContractTemplate) error {
	return r.db.Create(template).Error
}

func (r *repository) GetTemplateByID(id uint) (*entities.ContractTemplate, error) {
	var template entities.ContractTemplate
	err := r.db.First(&template, id).Error
	if err != nil {
		return nil, err
	}
	return &template, nil
}

func (r *repository) GetActiveTemplates() ([]entities.ContractTemplate, error) {
	var templates []entities.ContractTemplate
	err := r.db.Where("is_active = ?", true).Find(&templates).Error
	return templates, err
}

func (r *repository) UpdateTemplate(template *entities.ContractTemplate) error {
	return r.db.Save(template).Error
}