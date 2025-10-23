package entities

import (
	"time"
)

type Token struct {
	ID              uint      `json:"id" gorm:"primaryKey"`
	UserID          uint      `json:"user_id" gorm:"not null"`
	Name            string    `json:"name" gorm:"not null;size:100"`
	Symbol          string    `json:"symbol" gorm:"not null;size:20"`
	Description     string    `json:"description" gorm:"size:1000"`
	TotalSupply     string    `json:"total_supply" gorm:"not null"`
	Decimals        uint8     `json:"decimals" gorm:"default:18"`
	ContractAddress string    `json:"contract_address" gorm:"size:42"`
	Network         string    `json:"network" gorm:"not null;default:'ethereum'"`
	ImageURL        string    `json:"image_url" gorm:"size:500"`
	Website         string    `json:"website" gorm:"size:200"`
	Twitter         string    `json:"twitter" gorm:"size:200"`
	Telegram        string    `json:"telegram" gorm:"size:200"`
	Discord         string    `json:"discord" gorm:"size:200"`
	Status          string    `json:"status" gorm:"not null;default:'draft'"` // draft, deploying, deployed, failed
	DeploymentTxHash string   `json:"deployment_tx_hash" gorm:"size:66"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	
	// Relationships
	User User `json:"user" gorm:"foreignKey:UserID"`
}

type TokenFeatures struct {
	ID                uint   `json:"id" gorm:"primaryKey"`
	TokenID           uint   `json:"token_id" gorm:"not null"`
	IsMintable        bool   `json:"is_mintable" gorm:"default:false"`
	IsBurnable        bool   `json:"is_burnable" gorm:"default:false"`
	IsPausable        bool   `json:"is_pausable" gorm:"default:false"`
	HasMaxSupply      bool   `json:"has_max_supply" gorm:"default:false"`
	HasTaxes          bool   `json:"has_taxes" gorm:"default:false"`
	BuyTaxPercentage  uint8  `json:"buy_tax_percentage" gorm:"default:0"`
	SellTaxPercentage uint8  `json:"sell_tax_percentage" gorm:"default:0"`
	IsAntiWhale       bool   `json:"is_anti_whale" gorm:"default:false"`
	MaxTxAmount       string `json:"max_tx_amount"`
	MaxWalletAmount   string `json:"max_wallet_amount"`
	
	// Relationships
	Token Token `json:"token" gorm:"foreignKey:TokenID"`
}

type TokenAnalytics struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	TokenID      uint      `json:"token_id" gorm:"not null"`
	Holders      uint64    `json:"holders" gorm:"default:0"`
	Transactions uint64    `json:"transactions" gorm:"default:0"`
	Volume24h    string    `json:"volume_24h" gorm:"default:'0'"`
	MarketCap    string    `json:"market_cap" gorm:"default:'0'"`
	Price        string    `json:"price" gorm:"default:'0'"`
	UpdatedAt    time.Time `json:"updated_at"`
	
	// Relationships
	Token Token `json:"token" gorm:"foreignKey:TokenID"`
}