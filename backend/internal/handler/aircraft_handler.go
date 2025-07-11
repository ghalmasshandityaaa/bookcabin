package handler

import (
	"bookcabin-backend/internal/entity"
	"bookcabin-backend/internal/model"
	"bookcabin-backend/internal/usecase"
	"bookcabin-backend/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AircraftHandler struct {
	Log       *logrus.Logger
	UseCase   *usecase.AircraftUseCase
	Validator *validator.Validator
}

func NewAircraftHandler(
	log *logrus.Logger,
	validator *validator.Validator,
	useCase *usecase.AircraftUseCase,
) *AircraftHandler {
	return &AircraftHandler{
		Log:       log,
		UseCase:   useCase,
		Validator: validator,
	}
}

func (h *AircraftHandler) ListSeats(ctx *fiber.Ctx) error {
	method := "AircraftHandler.ListSeats"
	logger := h.Log.WithField("method", method)
	logger.Trace("BEGIN")

	request := &model.ListAircraftSeatsRequest{
		AircraftType: entity.AircraftType(ctx.Query("type", "")),
	}

	errValidation := h.Validator.ValidateStruct(request)
	if errValidation != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.WebResponse[any]{
			Ok:     false,
			Errors: errValidation,
		})
	}

	seats := h.UseCase.ListSeats(ctx.Context(), request)

	logger.Trace("END")

	return ctx.JSON(model.WebResponse[[]model.AircraftSeatResponse]{
		Ok:   true,
		Data: seats,
	})
}
