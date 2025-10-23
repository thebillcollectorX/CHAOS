package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tiagorlampert/CHAOS/presentation/http/request"
	"github.com/tiagorlampert/CHAOS/services/meme_coin"
)

type MemeCoinHandler struct {
	memeCoinService meme_coin.MemeCoinService
}

func NewMemeCoinHandler(memeCoinService meme_coin.MemeCoinService) *MemeCoinHandler {
	return &MemeCoinHandler{
		memeCoinService: memeCoinService,
	}
}

// CreateMemeCoin creates a new meme coin
func (h *MemeCoinHandler) CreateMemeCoin(c *gin.Context) {
	var req request.CreateMemeCoinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get user ID from context (assuming authentication middleware sets this)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	// Convert to service request
	serviceReq := &meme_coin.CreateMemeCoinRequest{
		Name:        req.Name,
		Symbol:      req.Symbol,
		Description: req.Description,
		TotalSupply: req.TotalSupply,
		Decimals:    req.Decimals,
		ImageURL:    req.ImageURL,
		Website:     req.Website,
		Twitter:     req.Twitter,
		Telegram:    req.Telegram,
		Discord:     req.Discord,
		Network:     req.Network,
		CreatorID:   userID.(string),
	}

	memeCoin, err := h.memeCoinService.CreateMemeCoin(serviceReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": memeCoin})
}

// GetMemeCoin retrieves a meme coin by ID
func (h *MemeCoinHandler) GetMemeCoin(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	memeCoin, err := h.memeCoinService.GetMemeCoinByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "meme coin not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": memeCoin})
}

// GetMemeCoinBySymbol retrieves a meme coin by symbol
func (h *MemeCoinHandler) GetMemeCoinBySymbol(c *gin.Context) {
	symbol := c.Param("symbol")
	if symbol == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "symbol is required"})
		return
	}

	memeCoin, err := h.memeCoinService.GetMemeCoinBySymbol(symbol)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "meme coin not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": memeCoin})
}

// GetMemeCoins retrieves all meme coins with pagination
func (h *MemeCoinHandler) GetMemeCoins(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 20
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	memeCoins, err := h.memeCoinService.GetAllMemeCoins(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve meme coins"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": memeCoins, "limit": limit, "offset": offset})
}

// SearchMemeCoins searches for meme coins
func (h *MemeCoinHandler) SearchMemeCoins(c *gin.Context) {
	var req request.SearchMemeCoinsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Limit == 0 {
		req.Limit = 20
	}

	memeCoins, err := h.memeCoinService.SearchMemeCoins(req.Query, req.Limit, req.Offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to search meme coins"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": memeCoins, "query": req.Query, "limit": req.Limit, "offset": req.Offset})
}

// GetUserMemeCoins retrieves meme coins created by the authenticated user
func (h *MemeCoinHandler) GetUserMemeCoins(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 || limit > 100 {
		limit = 20
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	memeCoins, err := h.memeCoinService.GetMemeCoinsByCreator(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve user meme coins"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": memeCoins, "limit": limit, "offset": offset})
}

// UpdateMemeCoin updates a meme coin
func (h *MemeCoinHandler) UpdateMemeCoin(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	// Get existing meme coin
	memeCoin, err := h.memeCoinService.GetMemeCoinByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "meme coin not found"})
		return
	}

	// Check if user owns the meme coin
	if memeCoin.CreatorID != userID.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized to update this meme coin"})
		return
	}

	var req request.UpdateMemeCoinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields
	if req.Name != "" {
		memeCoin.Name = req.Name
	}
	if req.Description != "" {
		memeCoin.Description = req.Description
	}
	if req.ImageURL != "" {
		memeCoin.ImageURL = req.ImageURL
	}
	if req.Website != "" {
		memeCoin.Website = req.Website
	}
	if req.Twitter != "" {
		memeCoin.Twitter = req.Twitter
	}
	if req.Telegram != "" {
		memeCoin.Telegram = req.Telegram
	}
	if req.Discord != "" {
		memeCoin.Discord = req.Discord
	}

	if err := h.memeCoinService.UpdateMemeCoin(memeCoin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update meme coin"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": memeCoin})
}

// DeployMemeCoin deploys a meme coin to the blockchain
func (h *MemeCoinHandler) DeployMemeCoin(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	// Get existing meme coin
	memeCoin, err := h.memeCoinService.GetMemeCoinByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "meme coin not found"})
		return
	}

	// Check if user owns the meme coin
	if memeCoin.CreatorID != userID.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized to deploy this meme coin"})
		return
	}

	deploymentTx, err := h.memeCoinService.DeployMemeCoin(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": deploymentTx})
}

// DeleteMemeCoin deletes a meme coin
func (h *MemeCoinHandler) DeleteMemeCoin(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	// Get existing meme coin
	memeCoin, err := h.memeCoinService.GetMemeCoinByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "meme coin not found"})
		return
	}

	// Check if user owns the meme coin
	if memeCoin.CreatorID != userID.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "not authorized to delete this meme coin"})
		return
	}

	// Check if meme coin is already deployed
	if memeCoin.Status == "deployed" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot delete deployed meme coin"})
		return
	}

	if err := h.memeCoinService.DeleteMemeCoin(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete meme coin"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "meme coin deleted successfully"})
}