package token

import (
	"github.com/tiagorlampert/CHAOS/entities"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(token *entities.Token) error {
	return r.db.Create(token).Error
}

func (r *repository) GetByID(id uint) (*entities.Token, error) {
	var token entities.Token
	err := r.db.Preload("User").Preload("TokenFeatures").First(&token, id).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *repository) GetByUserID(userID uint) ([]entities.Token, error) {
	var tokens []entities.Token
	err := r.db.Where("user_id = ?", userID).Preload("User").Find(&tokens).Error
	return tokens, err
}

func (r *repository) GetByContractAddress(address string) (*entities.Token, error) {
	var token entities.Token
	err := r.db.Where("contract_address = ?", address).Preload("User").First(&token).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *repository) Update(token *entities.Token) error {
	return r.db.Save(token).Error
}

func (r *repository) Delete(id uint) error {
	return r.db.Delete(&entities.Token{}, id).Error
}

func (r *repository) GetAll(limit, offset int) ([]entities.Token, error) {
	var tokens []entities.Token
	err := r.db.Limit(limit).Offset(offset).Preload("User").Find(&tokens).Error
	return tokens, err
}

func (r *repository) GetByStatus(status string) ([]entities.Token, error) {
	var tokens []entities.Token
	err := r.db.Where("status = ?", status).Preload("User").Find(&tokens).Error
	return tokens, err
}

func (r *repository) CreateFeatures(features *entities.TokenFeatures) error {
	return r.db.Create(features).Error
}

func (r *repository) UpdateFeatures(features *entities.TokenFeatures) error {
	return r.db.Save(features).Error
}

func (r *repository) GetFeaturesByTokenID(tokenID uint) (*entities.TokenFeatures, error) {
	var features entities.TokenFeatures
	err := r.db.Where("token_id = ?", tokenID).First(&features).Error
	if err != nil {
		return nil, err
	}
	return &features, nil
}

func (r *repository) UpdateAnalytics(analytics *entities.TokenAnalytics) error {
	return r.db.Save(analytics).Error
}

func (r *repository) GetAnalyticsByTokenID(tokenID uint) (*entities.TokenAnalytics, error) {
	var analytics entities.TokenAnalytics
	err := r.db.Where("token_id = ?", tokenID).First(&analytics).Error
	if err != nil {
		return nil, err
	}
	return &analytics, nil
}