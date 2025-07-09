package main

import (
	"bookcabin-backend/config"
	_ "bookcabin-backend/docs"
	"bookcabin-backend/internal/app"
	"bookcabin-backend/pkg/database/gorm"
	"bookcabin-backend/pkg/fiber"
	"bookcabin-backend/pkg/logger"
	"bookcabin-backend/pkg/middleware"
	"bookcabin-backend/pkg/validator"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// @title Voucher Seat Assignment API
// @version 1.0.0
// @description This API provides comprehensive voucher management and seat assignment functionality for flight booking systems. It enables checking existing voucher assignments, generating new vouchers with random seat allocation, and managing the entire voucher lifecycle.
// @description
// @description ## Features
// @description - Check existing voucher assignments for specific flights and dates
// @description - Generate vouchers with automatic random seat assignment
// @description - Prevent duplicate voucher assignments
// @description - RESTful API design with comprehensive error handling
// @description - Request validation and structured error responses
// @description
// @description ## Error Handling
// @description All endpoints return structured error responses with appropriate HTTP status codes:
// @description - 200: Success
// @description - 400: Bad Request (validation errors, malformed request body, invalid parameters)
// @description - 500: Internal Server Error
// @description
// @description ## Rate Limiting
// @description API requests are subject to rate limiting to ensure fair usage and system stability.
// @termsOfService http://swagger.io/terms/
// @contact.name Ghalmas Shanditya Putra Agung
// @contact.email ghalmas.shanditya.putra.agung@gmail.com
// @contact.url https://github.com/ghalmasshandityaaa
// @host localhost:3000
// @BasePath /api
// @schemes http https
func main() {
	conf := config.Read()

	log := logger.NewLogger(conf)
	log.Info("initialized logger")

	newValidator := validator.NewValidator()
	log.Info("initialized validator")

	// Connect to PostgresSQL under the GORM ORM
	db := gorm.NewGormDB(conf, log)
	log.Infof("database connected successfully")
	defer db.Close()

	// Initialize fiber application
	fiberApp := fiber.NewFiber(conf, log)
	log.Infof("initialized fiber server")

	// Setup middleware
	middleware.SetupMiddleware(fiberApp, conf)
	log.Infof("setup middleware for fiber server")

	// Bootstrap application
	app.Bootstrap(&app.BootstrapConfig{
		App:       fiberApp,
		DB:        db.DB(),
		Log:       log,
		Config:    conf,
		Validator: newValidator,
	})

	log.Info("setup exception middleware")
	middleware.SetupExceptionMiddleware(fiberApp)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	var serverShutdown sync.WaitGroup
	go func() {
		<-signalChan
		log.Info("Gracefully shutting down...")
		serverShutdown.Add(1)
		defer serverShutdown.Done()
		_ = fiberApp.ShutdownWithTimeout(60 * time.Second)
	}()

	log.Infof("starting server on port %d...", conf.App.Port)
	if err := fiberApp.Listen(fmt.Sprintf("%s:%d", conf.App.Host, conf.App.Port)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	serverShutdown.Wait()
	log.Info("Running cleanup tasks...")
}
