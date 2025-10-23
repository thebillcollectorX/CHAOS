package token

import (
	"github.com/tiagorlampert/CHAOS/entities"
)

type Repository interface {
	Insert(token entities.Token) error
	FindAll() ([]entities.Token, error)
	FindByID(id uint) (entities.Token, error)
	FindByUserID(userID uint) ([]entities.Token, error)
	Update(token entities.Token) error
	Delete(id uint) error
}
