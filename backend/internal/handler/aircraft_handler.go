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

// @Summary List seats for a given aircraft type
// @Description This endpoint returns a list of seats for a given aircraft type
// @Tags Aircrafts
// @Accept json
// @Produce json
// @Param type query string true "Aircraft type"
// @Success 200 {object} model.AircraftSeatSwaggerResponse "Successfully listed seats"
// @Failure 400 {object} model.ErrorResponse "Bad request - invalid payload, unprocessable entity, validation error"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /aircraft/seats [get]
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
