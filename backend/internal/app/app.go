package app

import (
	"bookcabin-backend/config"
	"bookcabin-backend/internal/handler"
	"bookcabin-backend/internal/repository"
	"bookcabin-backend/internal/route"
	"bookcabin-backend/internal/usecase"
	"bookcabin-backend/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	App       *fiber.App
	Log       *logrus.Logger
	Config    *config.Config
	DB        *gorm.DB
	Validator *validator.Validator
}

func Bootstrap(config *BootstrapConfig) {
	// init repositories
	voucherRepository := repository.NewVoucherRepository(config.Log)

	// init use cases
	voucherUseCase := usecase.NewVoucherUseCase(config.DB, config.Log, voucherRepository)
	aircraftUseCase := usecase.NewAircraftUseCase(config.DB, config.Log, voucherRepository)

	// init handlers
	voucherHandler := handler.NewVoucherHandler(config.Log, config.Validator, voucherUseCase)
	aircraftHandler := handler.NewAircraftHandler(config.Log, config.Validator, aircraftUseCase)

	// init routes
	appRoute := route.NewRoute(config.App, config.Log, voucherHandler, aircraftHandler)

	// setup routes
	appRoute.Setup()
}
