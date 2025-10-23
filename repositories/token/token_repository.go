package token

import (
	"github.com/tiagorlampert/CHAOS/entities"
)

type Repository interface {
	Create(token *entities.Token) error
	GetByID(id uint) (*entities.Token, error)
	GetByUserID(userID uint) ([]entities.Token, error)
	GetByContractAddress(address string) (*entities.Token, error)
	Update(token *entities.Token) error
	Delete(id uint) error
	GetAll(limit, offset int) ([]entities.Token, error)
	GetByStatus(status string) ([]entities.Token, error)
	CreateFeatures(features *entities.TokenFeatures) error
	UpdateFeatures(features *entities.TokenFeatures) error
	GetFeaturesByTokenID(tokenID uint) (*entities.TokenFeatures, error)
	UpdateAnalytics(analytics *entities.TokenAnalytics) error
	GetAnalyticsByTokenID(tokenID uint) (*entities.TokenAnalytics, error)
}