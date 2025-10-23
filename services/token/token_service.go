package token

import (
	"github.com/tiagorlampert/CHAOS/entities"
)

type Service interface {
	CreateToken(token *entities.Token, features *entities.TokenFeatures) error
	GetTokenByID(id uint) (*entities.Token, error)
	GetTokensByUserID(userID uint) ([]entities.Token, error)
	UpdateToken(token *entities.Token) error
	DeleteToken(id uint) error
	DeployToken(tokenID uint, networkID uint) (*entities.Transaction, error)
	GetTokenAnalytics(tokenID uint) (*entities.TokenAnalytics, error)
	UpdateTokenAnalytics(tokenID uint) error
	ValidateTokenData(token *entities.Token) error
	GenerateContractCode(token *entities.Token, features *entities.TokenFeatures) (string, error)
}