package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tiagorlampert/CHAOS/internal/environment"
	"github.com/tiagorlampert/CHAOS/internal/middleware"
	"github.com/tiagorlampert/CHAOS/presentation/http/handler"
	"github.com/tiagorlampert/CHAOS/services/meme_coin"
)

type MemeCoinController struct {
	Configuration  *environment.Configuration
	Logger         *logrus.Logger
	AuthMiddleware *middleware.JWT
	MemeCoinHandler *handler.MemeCoinHandler
}

func NewMemeCoinController(
	configuration *environment.Configuration,
	router *gin.Engine,
	log *logrus.Logger,
	authMiddleware *middleware.JWT,
	memeCoinService meme_coin.MemeCoinService,
) {
	controller := &MemeCoinController{
		Configuration:  configuration,
		Logger:         log,
		AuthMiddleware: authMiddleware,
		MemeCoinHandler: handler.NewMemeCoinHandler(memeCoinService),
	}

	// Public routes (no authentication required)
	publicGroup := router.Group("/meme-coins")
	{
		publicGroup.GET("/", controller.listMemeCoinsHandler)
		publicGroup.GET("/:id", controller.getMemeCoinHandler)
		publicGroup.GET("/symbol/:symbol", controller.getMemeCoinBySymbolHandler)
		publicGroup.GET("/search", controller.searchMemeCoinsHandler)
	}

	// Protected routes (authentication required)
	authGroup := router.Group("/meme-coins")
	authGroup.Use(authMiddleware.MiddlewareFunc())
	{
		authGroup.POST("/", controller.createMemeCoinHandler)
		authGroup.GET("/my", controller.getUserMemeCoinsHandler)
		authGroup.PUT("/:id", controller.updateMemeCoinHandler)
		authGroup.DELETE("/:id", controller.deleteMemeCoinHandler)
		authGroup.POST("/:id/deploy", controller.deployMemeCoinHandler)
		authGroup.POST("/preview-contract", controller.previewContractHandler)
	}

	// Web routes
	webGroup := router.Group("")
	{
		webGroup.GET("/meme-coins", controller.memeCoinListPageHandler)
		webGroup.GET("/meme-coins/create", controller.memeCoinCreatorPageHandler)
		webGroup.GET("/meme-coins/my", controller.memeCoinDashboardPageHandler)
		webGroup.GET("/meme-coins/:id", controller.memeCoinDetailPageHandler)
	}
}

// API Handlers
func (c *MemeCoinController) listMemeCoinsHandler(ctx *gin.Context) {
	c.MemeCoinHandler.GetMemeCoins(ctx)
}

func (c *MemeCoinController) getMemeCoinHandler(ctx *gin.Context) {
	c.MemeCoinHandler.GetMemeCoin(ctx)
}

func (c *MemeCoinController) getMemeCoinBySymbolHandler(ctx *gin.Context) {
	c.MemeCoinHandler.GetMemeCoinBySymbol(ctx)
}

func (c *MemeCoinController) searchMemeCoinsHandler(ctx *gin.Context) {
	c.MemeCoinHandler.SearchMemeCoins(ctx)
}

func (c *MemeCoinController) createMemeCoinHandler(ctx *gin.Context) {
	c.MemeCoinHandler.CreateMemeCoin(ctx)
}

func (c *MemeCoinController) getUserMemeCoinsHandler(ctx *gin.Context) {
	c.MemeCoinHandler.GetUserMemeCoins(ctx)
}

func (c *MemeCoinController) updateMemeCoinHandler(ctx *gin.Context) {
	c.MemeCoinHandler.UpdateMemeCoin(ctx)
}

func (c *MemeCoinController) deleteMemeCoinHandler(ctx *gin.Context) {
	c.MemeCoinHandler.DeleteMemeCoin(ctx)
}

func (c *MemeCoinController) deployMemeCoinHandler(ctx *gin.Context) {
	c.MemeCoinHandler.DeployMemeCoin(ctx)
}

func (c *MemeCoinController) previewContractHandler(ctx *gin.Context) {
	// This would generate contract preview
	ctx.JSON(200, gin.H{
		"contract_code": "// Contract preview would be generated here",
		"abi": "// ABI would be generated here",
		"bytecode": "// Bytecode would be generated here",
	})
}

// Web Page Handlers
func (c *MemeCoinController) memeCoinListPageHandler(ctx *gin.Context) {
	ctx.HTML(200, "meme_coin_list", gin.H{
		"title": "Meme Coins Marketplace",
	})
}

func (c *MemeCoinController) memeCoinCreatorPageHandler(ctx *gin.Context) {
	ctx.HTML(200, "meme_coin_creator", gin.H{
		"title": "Create Meme Coin",
	})
}

func (c *MemeCoinController) memeCoinDashboardPageHandler(ctx *gin.Context) {
	ctx.HTML(200, "meme_coin_dashboard", gin.H{
		"title": "My Meme Coins",
	})
}

func (c *MemeCoinController) memeCoinDetailPageHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.HTML(200, "meme_coin_detail", gin.H{
		"title": "Meme Coin Details",
		"id":    id,
	})
}