package token

import (
	"github.com/tiagorlampert/CHAOS/entities"
)

type Service interface {
	CreateToken(input CreateTokenInput) (CreateTokenOutput, error)
	GetAllTokens() ([]entities.Token, error)
	GetTokenByID(id uint) (entities.Token, error)
	GetTokensByUserID(userID uint) ([]entities.Token, error)
	UpdateTokenStatus(id uint, status string, address string, txHash string) error
}

type CreateTokenInput struct {
	Name        string
	Symbol      string
	TotalSupply string
	Description string
	ImageURL    string
	Network     string
	Decimals    int
	UserID      uint
}

type CreateTokenOutput struct {
	ID      uint   `json:"id"`
	Address string `json:"address,omitempty"`
	TxHash  string `json:"tx_hash,omitempty"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
