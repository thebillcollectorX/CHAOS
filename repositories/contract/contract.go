package contract

import "github.com/tiagorlampert/CHAOS/entities"

type ContractRepository interface {
	Create(contract *entities.TokenContract) error
	GetByMemeCoinID(memeCoinID string) (*entities.TokenContract, error)
	Update(contract *entities.TokenContract) error
	Delete(id string) error
}