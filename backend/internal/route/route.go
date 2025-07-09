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
}

func NewRoute(
	app *fiber.App,
	logger *logrus.Logger,
	voucherHandler *handler.VoucherHandler,
) *Route {
	return &Route{
		App:            app,
		Log:            logger,
		VoucherHandler: voucherHandler,
	}
}

func (a *Route) Setup() {
	a.Log.Info("setting up routes")

	a.SetupVoucherRoute()
	a.SetupSwaggerRoute()
}
