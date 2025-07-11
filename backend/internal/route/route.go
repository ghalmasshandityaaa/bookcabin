package route

import (
	"bookcabin-backend/internal/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Route struct {
	App            *fiber.App
	Log            *logrus.Logger
	VoucherHandler *handler.VoucherHandler
	AircraftHander *handler.AircraftHandler
}

func NewRoute(
	app *fiber.App,
	logger *logrus.Logger,
	voucherHandler *handler.VoucherHandler,
	aircraftHandler *handler.AircraftHandler,
) *Route {
	return &Route{
		App:            app,
		Log:            logger,
		VoucherHandler: voucherHandler,
		AircraftHander: aircraftHandler,
	}
}

func (a *Route) Setup() {
	a.Log.Info("setting up routes")

	a.SetupVoucherRoute()
	a.SetupAircraftRoute()
	a.SetupSwaggerRoute()
}
