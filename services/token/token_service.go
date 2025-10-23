package token

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/tiagorlampert/CHAOS/entities"
	"github.com/tiagorlampert/CHAOS/repositories/token"
	"strconv"
	"time"
)

type service struct {
	repository token.Repository
}

func NewService(repository token.Repository) Service {
	return &service{repository: repository}
}

func (s *service) CreateToken(input CreateTokenInput) (CreateTokenOutput, error) {
	// Validate input
	if input.Name == "" || input.Symbol == "" || input.TotalSupply == "" || input.Network == "" {
		return CreateTokenOutput{}, errors.New("missing required fields")
	}

	// Validate total supply is a valid number
	supply, err := strconv.ParseInt(input.TotalSupply, 10, 64)
	if err != nil || supply <= 0 {
		return CreateTokenOutput{}, errors.New("invalid total supply")
	}

	// Create token entity
	newToken := entities.Token{
		Name:        input.Name,
		Symbol:      input.Symbol,
		TotalSupply: input.TotalSupply,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		Network:     input.Network,
		Decimals:    input.Decimals,
		UserID:      input.UserID,
		Status:      "pending",
	}

	// Insert into database
	if err := s.repository.Insert(newToken); err != nil {
		return CreateTokenOutput{}, fmt.Errorf("failed to create token: %w", err)
	}

	// Simulate blockchain deployment (in a real scenario, this would interact with actual blockchain)
	// For demo purposes, we'll generate a mock address and transaction hash
	go s.simulateBlockchainDeployment(newToken.ID)

	return CreateTokenOutput{
		ID:      newToken.ID,
		Status:  "pending",
		Message: "Token creation initiated. Deployment in progress...",
	}, nil
}

func (s *service) simulateBlockchainDeployment(tokenID uint) {
	// Simulate deployment delay (2-5 seconds)
	time.Sleep(3 * time.Second)

	// Generate mock blockchain address and transaction hash
	address := s.generateMockAddress()
	txHash := s.generateMockTxHash()

	// Update token status
	token, err := s.repository.FindByID(tokenID)
	if err != nil {
		return
	}

	token.Status = "deployed"
	token.Address = address
	token.TxHash = txHash

	s.repository.Update(token)
}

func (s *service) generateMockAddress() string {
	// Generate a mock Ethereum-style address
	bytes := make([]byte, 20)
	rand.Read(bytes)
	return "0x" + hex.EncodeToString(bytes)
}

func (s *service) generateMockTxHash() string {
	// Generate a mock transaction hash
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return "0x" + hex.EncodeToString(bytes)
}

func (s *service) GetAllTokens() ([]entities.Token, error) {
	return s.repository.FindAll()
}

func (s *service) GetTokenByID(id uint) (entities.Token, error) {
	return s.repository.FindByID(id)
}

func (s *service) GetTokensByUserID(userID uint) ([]entities.Token, error) {
	return s.repository.FindByUserID(userID)
}

func (s *service) UpdateTokenStatus(id uint, status string, address string, txHash string) error {
	token, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}

	token.Status = status
	if address != "" {
		token.Address = address
	}
	if txHash != "" {
		token.TxHash = txHash
	}

	return s.repository.Update(token)
}
