package token

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/tiagorlampert/CHAOS/entities"
	blockchainRepo "github.com/tiagorlampert/CHAOS/repositories/blockchain"
	tokenRepo "github.com/tiagorlampert/CHAOS/repositories/token"
)

type service struct {
	logger          *logrus.Logger
	tokenRepo       tokenRepo.Repository
	blockchainRepo  blockchainRepo.Repository
}

func NewTokenService(logger *logrus.Logger, tokenRepo tokenRepo.Repository, blockchainRepo blockchainRepo.Repository) Service {
	return &service{
		logger:         logger,
		tokenRepo:      tokenRepo,
		blockchainRepo: blockchainRepo,
	}
}

func (s *service) CreateToken(token *entities.Token, features *entities.TokenFeatures) error {
	// Validate token data
	if err := s.ValidateTokenData(token); err != nil {
		return err
	}

	// Create token
	if err := s.tokenRepo.Create(token); err != nil {
		s.logger.WithError(err).Error("Failed to create token")
		return err
	}

	// Create token features
	features.TokenID = token.ID
	if err := s.tokenRepo.CreateFeatures(features); err != nil {
		s.logger.WithError(err).Error("Failed to create token features")
		return err
	}

	// Create initial analytics
	analytics := &entities.TokenAnalytics{
		TokenID:      token.ID,
		Holders:      0,
		Transactions: 0,
		Volume24h:    "0",
		MarketCap:    "0",
		Price:        "0",
		UpdatedAt:    time.Now(),
	}
	
	if err := s.tokenRepo.UpdateAnalytics(analytics); err != nil {
		s.logger.WithError(err).Error("Failed to create token analytics")
		return err
	}

	return nil
}

func (s *service) GetTokenByID(id uint) (*entities.Token, error) {
	return s.tokenRepo.GetByID(id)
}

func (s *service) GetTokensByUserID(userID uint) ([]entities.Token, error) {
	return s.tokenRepo.GetByUserID(userID)
}

func (s *service) UpdateToken(token *entities.Token) error {
	if err := s.ValidateTokenData(token); err != nil {
		return err
	}
	return s.tokenRepo.Update(token)
}

func (s *service) DeleteToken(id uint) error {
	return s.tokenRepo.Delete(id)
}

func (s *service) DeployToken(tokenID uint, networkID uint) (*entities.Transaction, error) {
	token, err := s.tokenRepo.GetByID(tokenID)
	if err != nil {
		return nil, err
	}

	if token.Status != "draft" {
		return nil, errors.New("token is not in draft status")
	}

	// Update token status to deploying
	token.Status = "deploying"
	if err := s.tokenRepo.Update(token); err != nil {
		return nil, err
	}

	// Create deployment transaction record
	tx := &entities.Transaction{
		TokenID:   tokenID,
		UserID:    token.UserID,
		Type:      "deploy",
		Status:    "pending",
		NetworkID: networkID,
	}

	if err := s.blockchainRepo.CreateTransaction(tx); err != nil {
		return nil, err
	}

	// TODO: Implement actual blockchain deployment logic here
	// This would involve:
	// 1. Generating contract bytecode
	// 2. Estimating gas
	// 3. Sending deployment transaction
	// 4. Monitoring transaction status

	return tx, nil
}

func (s *service) GetTokenAnalytics(tokenID uint) (*entities.TokenAnalytics, error) {
	return s.tokenRepo.GetAnalyticsByTokenID(tokenID)
}

func (s *service) UpdateTokenAnalytics(tokenID uint) error {
	// TODO: Implement analytics update logic
	// This would fetch data from blockchain/DEX APIs
	return nil
}

func (s *service) ValidateTokenData(token *entities.Token) error {
	if token.Name == "" {
		return errors.New("token name is required")
	}

	if token.Symbol == "" {
		return errors.New("token symbol is required")
	}

	// Validate symbol format (alphanumeric, 2-20 characters)
	symbolRegex := regexp.MustCompile(`^[A-Za-z0-9]{2,20}$`)
	if !symbolRegex.MatchString(token.Symbol) {
		return errors.New("token symbol must be 2-20 alphanumeric characters")
	}

	if token.TotalSupply == "" {
		return errors.New("total supply is required")
	}

	// Validate total supply is a valid number
	if _, err := strconv.ParseFloat(token.TotalSupply, 64); err != nil {
		return errors.New("total supply must be a valid number")
	}

	// Validate decimals (0-18)
	if token.Decimals > 18 {
		return errors.New("decimals cannot exceed 18")
	}

	return nil
}

func (s *service) GenerateContractCode(token *entities.Token, features *entities.TokenFeatures) (string, error) {
	// Basic ERC-20 template
	template := `// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract %s is ERC20, Ownable {
    constructor() ERC20("%s", "%s") {
        _mint(msg.sender, %s * 10**decimals());
    }
}`

	// Calculate total supply with decimals
	totalSupply, _ := strconv.ParseFloat(token.TotalSupply, 64)
	totalSupplyStr := fmt.Sprintf("%.0f", totalSupply)

	// Generate contract name from symbol
	contractName := strings.ToUpper(token.Symbol) + "Token"

	contractCode := fmt.Sprintf(template, contractName, token.Name, token.Symbol, totalSupplyStr)

	// Add additional features if enabled
	if features.IsMintable {
		contractCode = strings.Replace(contractCode, "contract "+contractName+" is ERC20, Ownable {", 
			"contract "+contractName+" is ERC20, Ownable {\n    function mint(address to, uint256 amount) public onlyOwner {\n        _mint(to, amount);\n    }", 1)
	}

	if features.IsBurnable {
		contractCode = strings.Replace(contractCode, "import \"@openzeppelin/contracts/access/Ownable.sol\";", 
			"import \"@openzeppelin/contracts/access/Ownable.sol\";\nimport \"@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol\";", 1)
		contractCode = strings.Replace(contractCode, "contract "+contractName+" is ERC20, Ownable {", 
			"contract "+contractName+" is ERC20, ERC20Burnable, Ownable {", 1)
	}

	if features.IsPausable {
		contractCode = strings.Replace(contractCode, "import \"@openzeppelin/contracts/access/Ownable.sol\";", 
			"import \"@openzeppelin/contracts/access/Ownable.sol\";\nimport \"@openzeppelin/contracts/security/Pausable.sol\";", 1)
		contractCode = strings.Replace(contractCode, "contract "+contractName+" is ERC20, Ownable {", 
			"contract "+contractName+" is ERC20, Pausable, Ownable {", 1)
	}

	return contractCode, nil
}