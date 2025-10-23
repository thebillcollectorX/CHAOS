package entities

import "time"

type MemeCoin struct {
	DBModel
	Name            string    `json:"name" binding:"required" gorm:"uniqueIndex"`
	Symbol          string    `json:"symbol" binding:"required" gorm:"uniqueIndex"`
	Description     string    `json:"description"`
	TotalSupply     string    `json:"total_supply" binding:"required"`
	Decimals        uint8     `json:"decimals" binding:"required"`
	ImageURL        string    `json:"image_url"`
	Website         string    `json:"website"`
	Twitter         string    `json:"twitter"`
	Telegram        string    `json:"telegram"`
	Discord         string    `json:"discord"`
	ContractAddress string    `json:"contract_address"`
	Network         string    `json:"network" binding:"required"` // ethereum, bsc, polygon, etc.
	Status          string    `json:"status"`                     // pending, deployed, failed
	DeploymentHash  string    `json:"deployment_hash"`
	DeployedAt      *time.Time `json:"deployed_at"`
	CreatorID       string    `json:"creator_id" binding:"required"`
	Creator         User      `json:"creator" gorm:"foreignKey:CreatorID"`
	Price           float64   `json:"price"` // Price in ETH/BNB for deployment
	IsVerified      bool      `json:"is_verified" gorm:"default:false"`
}

type TokenContract struct {
	DBModel
	MemeCoinID      string `json:"meme_coin_id" binding:"required"`
	MemeCoin        MemeCoin `json:"meme_coin" gorm:"foreignKey:MemeCoinID"`
	ContractCode    string `json:"contract_code"`
	ABI             string `json:"abi"`
	Bytecode        string `json:"bytecode"`
	ConstructorArgs string `json:"constructor_args"`
	Network         string `json:"network" binding:"required"`
	GasLimit        uint64 `json:"gas_limit"`
	GasPrice        string `json:"gas_price"`
}

type DeploymentTransaction struct {
	DBModel
	MemeCoinID      string    `json:"meme_coin_id" binding:"required"`
	MemeCoin        MemeCoin  `json:"meme_coin" gorm:"foreignKey:MemeCoinID"`
	TransactionHash string    `json:"transaction_hash" binding:"required"`
	BlockNumber     uint64    `json:"block_number"`
	GasUsed         uint64    `json:"gas_used"`
	GasPrice        string    `json:"gas_price"`
	Status          string    `json:"status"` // pending, confirmed, failed
	DeployedAt      time.Time `json:"deployed_at"`
	Network         string    `json:"network" binding:"required"`
}

type Payment struct {
	DBModel
	MemeCoinID      string    `json:"meme_coin_id" binding:"required"`
	MemeCoin        MemeCoin  `json:"meme_coin" gorm:"foreignKey:MemeCoinID"`
	Amount          float64   `json:"amount" binding:"required"`
	Currency        string    `json:"currency" binding:"required"` // ETH, BNB, USDT, etc.
	PaymentMethod   string    `json:"payment_method"`              // crypto, stripe, paypal
	TransactionHash string    `json:"transaction_hash"`
	Status          string    `json:"status"` // pending, completed, failed, refunded
	PaidAt          *time.Time `json:"paid_at"`
	UserID          string    `json:"user_id" binding:"required"`
	User            User      `json:"user" gorm:"foreignKey:UserID"`
}