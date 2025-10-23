package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tiagorlampert/CHAOS/entities"
	"github.com/tiagorlampert/CHAOS/presentation/http/request"
	"github.com/tiagorlampert/CHAOS/services/token"
)

type TokenController struct {
	logger       *logrus.Logger
	tokenService token.Service
}

func NewTokenController(logger *logrus.Logger, tokenService token.Service) *TokenController {
	return &TokenController{
		logger:       logger,
		tokenService: tokenService,
	}
}

func (tc *TokenController) CreateToken(c *gin.Context) {
	var req request.CreateTokenRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from JWT token
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	token := &entities.Token{
		UserID:      userID.(uint),
		Name:        req.Name,
		Symbol:      req.Symbol,
		Description: req.Description,
		TotalSupply: req.TotalSupply,
		Decimals:    req.Decimals,
		Network:     req.Network,
		ImageURL:    req.ImageURL,
		Website:     req.Website,
		Twitter:     req.Twitter,
		Telegram:    req.Telegram,
		Discord:     req.Discord,
		Status:      "draft",
	}

	features := &entities.TokenFeatures{
		IsMintable:        req.IsMintable,
		IsBurnable:        req.IsBurnable,
		IsPausable:        req.IsPausable,
		HasMaxSupply:      req.HasMaxSupply,
		HasTaxes:          req.HasTaxes,
		BuyTaxPercentage:  req.BuyTaxPercentage,
		SellTaxPercentage: req.SellTaxPercentage,
		IsAntiWhale:       req.IsAntiWhale,
		MaxTxAmount:       req.MaxTxAmount,
		MaxWalletAmount:   req.MaxWalletAmount,
	}

	if err := tc.tokenService.CreateToken(token, features); err != nil {
		tc.logger.WithError(err).Error("Failed to create token")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Token created successfully",
		"token":   token,
	})
}

func (tc *TokenController) GetToken(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token ID"})
		return
	}

	token, err := tc.tokenService.GetTokenByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Token not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (tc *TokenController) GetUserTokens(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	tokens, err := tc.tokenService.GetTokensByUserID(userID.(uint))
	if err != nil {
		tc.logger.WithError(err).Error("Failed to get user tokens")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get tokens"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tokens": tokens})
}

func (tc *TokenController) UpdateToken(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token ID"})
		return
	}

	var req request.UpdateTokenRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := tc.tokenService.GetTokenByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Token not found"})
		return
	}

	// Verify ownership
	userID, _ := c.Get("userID")
	if token.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Update fields
	if req.Name != "" {
		token.Name = req.Name
	}
	if req.Description != "" {
		token.Description = req.Description
	}
	if req.ImageURL != "" {
		token.ImageURL = req.ImageURL
	}
	if req.Website != "" {
		token.Website = req.Website
	}
	if req.Twitter != "" {
		token.Twitter = req.Twitter
	}
	if req.Telegram != "" {
		token.Telegram = req.Telegram
	}
	if req.Discord != "" {
		token.Discord = req.Discord
	}

	if err := tc.tokenService.UpdateToken(token); err != nil {
		tc.logger.WithError(err).Error("Failed to update token")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Token updated successfully",
		"token":   token,
	})
}

func (tc *TokenController) DeployToken(c *gin.Context) {
	var req request.DeployTokenRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verify token ownership
	token, err := tc.tokenService.GetTokenByID(req.TokenID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Token not found"})
		return
	}

	userID, _ := c.Get("userID")
	if token.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	tx, err := tc.tokenService.DeployToken(req.TokenID, req.NetworkID)
	if err != nil {
		tc.logger.WithError(err).Error("Failed to deploy token")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Token deployment initiated",
		"transaction": tx,
	})
}

func (tc *TokenController) GetTokenAnalytics(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token ID"})
		return
	}

	analytics, err := tc.tokenService.GetTokenAnalytics(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Analytics not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"analytics": analytics})
}

func (tc *TokenController) DeleteToken(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token ID"})
		return
	}

	// Verify token ownership
	token, err := tc.tokenService.GetTokenByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Token not found"})
		return
	}

	userID, _ := c.Get("userID")
	if token.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	// Only allow deletion of draft tokens
	if token.Status != "draft" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot delete deployed tokens"})
		return
	}

	if err := tc.tokenService.DeleteToken(uint(id)); err != nil {
		tc.logger.WithError(err).Error("Failed to delete token")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token deleted successfully"})
}