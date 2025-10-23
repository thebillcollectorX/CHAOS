package meme_coin

import (
	"github.com/tiagorlampert/CHAOS/entities"
	"gorm.io/gorm"
)

type MemeCoinRepository interface {
	Create(memeCoin *entities.MemeCoin) error
	GetByID(id string) (*entities.MemeCoin, error)
	GetBySymbol(symbol string) (*entities.MemeCoin, error)
	GetByCreator(creatorID string) ([]entities.MemeCoin, error)
	GetAll(limit, offset int) ([]entities.MemeCoin, error)
	Update(memeCoin *entities.MemeCoin) error
	Delete(id string) error
	Search(query string, limit, offset int) ([]entities.MemeCoin, error)
	GetByStatus(status string) ([]entities.MemeCoin, error)
	GetByNetwork(network string) ([]entities.MemeCoin, error)
}

type memeCoinRepository struct {
	db *gorm.DB
}

func NewMemeCoinRepository(db *gorm.DB) MemeCoinRepository {
	return &memeCoinRepository{db: db}
}

func (r *memeCoinRepository) Create(memeCoin *entities.MemeCoin) error {
	return r.db.Create(memeCoin).Error
}

func (r *memeCoinRepository) GetByID(id string) (*entities.MemeCoin, error) {
	var memeCoin entities.MemeCoin
	err := r.db.Preload("Creator").First(&memeCoin, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &memeCoin, nil
}

func (r *memeCoinRepository) GetBySymbol(symbol string) (*entities.MemeCoin, error) {
	var memeCoin entities.MemeCoin
	err := r.db.Preload("Creator").First(&memeCoin, "symbol = ?", symbol).Error
	if err != nil {
		return nil, err
	}
	return &memeCoin, nil
}

func (r *memeCoinRepository) GetByCreator(creatorID string) ([]entities.MemeCoin, error) {
	var memeCoins []entities.MemeCoin
	err := r.db.Preload("Creator").Where("creator_id = ?", creatorID).Find(&memeCoins).Error
	return memeCoins, err
}

func (r *memeCoinRepository) GetAll(limit, offset int) ([]entities.MemeCoin, error) {
	var memeCoins []entities.MemeCoin
	query := r.db.Preload("Creator")
	
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}
	
	err := query.Order("created_at DESC").Find(&memeCoins).Error
	return memeCoins, err
}

func (r *memeCoinRepository) Update(memeCoin *entities.MemeCoin) error {
	return r.db.Save(memeCoin).Error
}

func (r *memeCoinRepository) Delete(id string) error {
	return r.db.Delete(&entities.MemeCoin{}, "id = ?", id).Error
}

func (r *memeCoinRepository) Search(query string, limit, offset int) ([]entities.MemeCoin, error) {
	var memeCoins []entities.MemeCoin
	dbQuery := r.db.Preload("Creator").Where(
		"name ILIKE ? OR symbol ILIKE ? OR description ILIKE ?",
		"%"+query+"%", "%"+query+"%", "%"+query+"%",
	)
	
	if limit > 0 {
		dbQuery = dbQuery.Limit(limit)
	}
	if offset > 0 {
		dbQuery = dbQuery.Offset(offset)
	}
	
	err := dbQuery.Order("created_at DESC").Find(&memeCoins).Error
	return memeCoins, err
}

func (r *memeCoinRepository) GetByStatus(status string) ([]entities.MemeCoin, error) {
	var memeCoins []entities.MemeCoin
	err := r.db.Preload("Creator").Where("status = ?", status).Order("created_at DESC").Find(&memeCoins).Error
	return memeCoins, err
}

func (r *memeCoinRepository) GetByNetwork(network string) ([]entities.MemeCoin, error) {
	var memeCoins []entities.MemeCoin
	err := r.db.Preload("Creator").Where("network = ?", network).Order("created_at DESC").Find(&memeCoins).Error
	return memeCoins, err
}