package main

import (
	"bookcabin-backend/config"
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
		_ = <-signalChan
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
