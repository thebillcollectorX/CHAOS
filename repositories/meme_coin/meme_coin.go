package meme_coin

import "github.com/tiagorlampert/CHAOS/entities"

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