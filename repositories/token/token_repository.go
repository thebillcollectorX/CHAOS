package token

import (
	"errors"
	"gorm.io/gorm"
	"github.com/tiagorlampert/CHAOS/entities"
	"github.com/tiagorlampert/CHAOS/repositories"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Insert(token entities.Token) error {
	if err := r.db.Create(&token).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) FindAll() ([]entities.Token, error) {
	var tokens []entities.Token
	if err := r.db.Order("created_at DESC").Find(&tokens).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entities.Token{}, nil
		}
		return nil, err
	}
	return tokens, nil
}

func (r *repository) FindByID(id uint) (entities.Token, error) {
	var token entities.Token
	if err := r.db.Where("id = ?", id).First(&token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Token{}, repositories.ErrNotFound
		}
		return entities.Token{}, err
	}
	return token, nil
}

func (r *repository) FindByUserID(userID uint) ([]entities.Token, error) {
	var tokens []entities.Token
	if err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&tokens).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []entities.Token{}, nil
		}
		return nil, err
	}
	return tokens, nil
}

func (r *repository) Update(token entities.Token) error {
	if err := r.db.Save(&token).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) Delete(id uint) error {
	if err := r.db.Where("id = ?", id).Delete(&entities.Token{}).Error; err != nil {
		return err
	}
	return nil
}
