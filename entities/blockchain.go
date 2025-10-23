package entities

import (
	"time"
)

type Network struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null;unique;size:50"`
	DisplayName string `json:"display_name" gorm:"not null;size:100"`
	ChainID     uint64 `json:"chain_id" gorm:"not null;unique"`
	RpcURL      string `json:"rpc_url" gorm:"not null;size:500"`
	ExplorerURL string `json:"explorer_url" gorm:"size:500"`
	Currency    string `json:"currency" gorm:"not null;size:10"`
	IsActive    bool   `json:"is_active" gorm:"default:true"`
	GasPrice    string `json:"gas_price" gorm:"default:'20000000000'"` // in wei
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Wallet struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	Address   string    `json:"address" gorm:"not null;size:42"`
	Type      string    `json:"type" gorm:"not null;size:20"` // metamask, walletconnect, etc.
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	// Relationships
	User User `json:"user" gorm:"foreignKey:UserID"`
}

type Transaction struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	TokenID     uint      `json:"token_id"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	Hash        string    `json:"hash" gorm:"not null;unique;size:66"`
	Type        string    `json:"type" gorm:"not null;size:20"` // deploy, transfer, mint, burn
	Status      string    `json:"status" gorm:"not null;default:'pending'"` // pending, confirmed, failed
	GasUsed     string    `json:"gas_used"`
	GasPrice    string    `json:"gas_price"`
	BlockNumber uint64    `json:"block_number"`
	NetworkID   uint      `json:"network_id" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	
	// Relationships
	Token   *Token   `json:"token,omitempty" gorm:"foreignKey:TokenID"`
	User    User     `json:"user" gorm:"foreignKey:UserID"`
	Network Network  `json:"network" gorm:"foreignKey:NetworkID"`
}

type ContractTemplate struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null;unique;size:100"`
	Description string    `json:"description" gorm:"size:500"`
	Version     string    `json:"version" gorm:"not null;size:20"`
	Solidity    string    `json:"solidity" gorm:"type:text;not null"`
	Bytecode    string    `json:"bytecode" gorm:"type:text"`
	ABI         string    `json:"abi" gorm:"type:text"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}