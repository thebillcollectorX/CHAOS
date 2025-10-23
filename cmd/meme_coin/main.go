package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/tiagorlampert/CHAOS/entities"
	"github.com/tiagorlampert/CHAOS/infrastructure/database"
	"github.com/tiagorlampert/CHAOS/internal/environment"
	"github.com/tiagorlampert/CHAOS/internal/middleware"
	"github.com/tiagorlampert/CHAOS/presentation/http"
	"github.com/tiagorlampert/CHAOS/presentation/http/controller"
	"github.com/tiagorlampert/CHAOS/repositories/auth"
	"github.com/tiagorlampert/CHAOS/repositories/contract"
	"github.com/tiagorlampert/CHAOS/repositories/device"
	"github.com/tiagorlampert/CHAOS/repositories/meme_coin"
	"github.com/tiagorlampert/CHAOS/repositories/user"
	"github.com/tiagorlampert/CHAOS/services/auth"
	"github.com/tiagorlampert/CHAOS/services/client"
	"github.com/tiagorlampert/CHAOS/services/device"
	"github.com/tiagorlampert/CHAOS/services/meme_coin"
	"github.com/tiagorlampert/CHAOS/services/url"
	"github.com/tiagorlampert/CHAOS/services/user"
)

func main() {
	// Load configuration
	config := environment.Load()

	// Initialize logger
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	// Initialize database
	db, err := database.NewProvider(config.Database)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate database
	err = db.AutoMigrate(
		&entities.User{},
		&entities.Device{},
		&entities.MemeCoin{},
		&entities.TokenContract{},
		&entities.DeploymentTransaction{},
		&entities.Payment{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize repositories
	userRepo := user.NewUserRepository(db)
	authRepo := auth.NewAuthRepository(db)
	deviceRepo := device.NewDeviceRepository(db)
	memeCoinRepo := meme_coin.NewMemeCoinRepository(db)
	contractRepo := contract.NewContractRepository(db)

	// Initialize services
	userService := user.NewUserService(userRepo)
	authService := auth.NewAuthService(authRepo, userRepo)
	deviceService := device.NewDeviceService(deviceRepo)
	memeCoinService := meme_coin.NewMemeCoinService(memeCoinRepo, contractRepo)
	clientService := client.NewClientService()
	urlService := url.NewUrlService()

	// Initialize JWT middleware
	jwtMiddleware := middleware.NewJWT(authService, logger)

	// Initialize HTTP router
	router := http.NewRouter()

	// Initialize controllers
	controller.NewMemeCoinController(config, router, logger, jwtMiddleware, memeCoinService)

	// Start server
	logger.Info("Starting Meme Coin Creator server on port ", config.Server.Port)
	if err := http.NewServer(router, config); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}