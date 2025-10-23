package contract

import (
	"github.com/tiagorlampert/CHAOS/entities"
	"gorm.io/gorm"
)

type ContractRepository interface {
	Create(contract *entities.TokenContract) error
	GetByMemeCoinID(memeCoinID string) (*entities.TokenContract, error)
	Update(contract *entities.TokenContract) error
	Delete(id string) error
}

type contractRepository struct {
	db *gorm.DB
}

func NewContractRepository(db *gorm.DB) ContractRepository {
	return &contractRepository{db: db}
}

func (r *contractRepository) Create(contract *entities.TokenContract) error {
	return r.db.Create(contract).Error
}

func (r *contractRepository) GetByMemeCoinID(memeCoinID string) (*entities.TokenContract, error) {
	var contract entities.TokenContract
	err := r.db.Preload("MemeCoin").First(&contract, "meme_coin_id = ?", memeCoinID).Error
	if err != nil {
		return nil, err
	}
	return &contract, nil
}

func (r *contractRepository) Update(contract *entities.TokenContract) error {
	return r.db.Save(contract).Error
}

func (r *contractRepository) Delete(id string) error {
	return r.db.Delete(&entities.TokenContract{}, "id = ?", id).Error
}