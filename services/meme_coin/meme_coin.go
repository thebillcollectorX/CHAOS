package meme_coin

import "github.com/tiagorlampert/CHAOS/entities"

type MemeCoinService interface {
	CreateMemeCoin(req *CreateMemeCoinRequest) (*entities.MemeCoin, error)
	GetMemeCoinByID(id string) (*entities.MemeCoin, error)
	GetMemeCoinBySymbol(symbol string) (*entities.MemeCoin, error)
	GetMemeCoinsByCreator(creatorID string) ([]entities.MemeCoin, error)
	GetAllMemeCoins(limit, offset int) ([]entities.MemeCoin, error)
	SearchMemeCoins(query string, limit, offset int) ([]entities.MemeCoin, error)
	UpdateMemeCoin(memeCoin *entities.MemeCoin) error
	DeleteMemeCoin(id string) error
	DeployMemeCoin(id string) (*entities.DeploymentTransaction, error)
	GenerateContractCode(memeCoin *entities.MemeCoin) (string, string, string, error)
	ValidateMemeCoinData(req *CreateMemeCoinRequest) error
	CalculateDeploymentCost(network string) (float64, error)
}

type CreateMemeCoinRequest struct {
	Name        string `json:"name" binding:"required"`
	Symbol      string `json:"symbol" binding:"required"`
	Description string `json:"description"`
	TotalSupply string `json:"total_supply" binding:"required"`
	Decimals    uint8  `json:"decimals" binding:"required"`
	ImageURL    string `json:"image_url"`
	Website     string `json:"website"`
	Twitter     string `json:"twitter"`
	Telegram    string `json:"telegram"`
	Discord     string `json:"discord"`
	Network     string `json:"network" binding:"required"`
	CreatorID   string `json:"creator_id" binding:"required"`
}