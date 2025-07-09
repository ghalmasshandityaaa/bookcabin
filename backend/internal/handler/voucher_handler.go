package handler

import (
	"bookcabin-backend/internal/model"
	"bookcabin-backend/internal/usecase"
	"bookcabin-backend/pkg/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type VoucherHandler struct {
	Log       *logrus.Logger
	UseCase   *usecase.VoucherUseCase
	Validator *validator.Validator
}

func NewVoucherHandler(
	log *logrus.Logger,
	validator *validator.Validator,
	useCase *usecase.VoucherUseCase,
) *VoucherHandler {
	return &VoucherHandler{
		Log:       log,
		UseCase:   useCase,
		Validator: validator,
	}
}

// Check verifies if a flight already has assigned vouchers for the specified date
// @Summary Check if flight has assigned vouchers
// @Description This endpoint checks whether a specific flight already has vouchers assigned for a given date. It helps prevent duplicate voucher assignments and ensures proper voucher management.
// @Tags Vouchers
// @Accept json
// @Produce json
// @Param request body model.CheckVoucherRequest true "Request body containing flight information and date to check for existing vouchers"
// @Success 200 {object} model.CheckVoucherResponse "Successfully checked voucher existence status"
// @Failure 400 {object} model.ErrorResponse "Bad request - invalid payload, unprocessable entity or validation error"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /check [post]
func (v *VoucherHandler) Check(ctx *fiber.Ctx) error {
	method := "VoucherHandler.Check"
	logger := v.Log.WithField("method", method)
	logger.Trace("BEGIN")

	request := new(model.CheckVoucherRequest)
	if err := ctx.BodyParser(request); err != nil {
		logger.Error("failed parse body: ", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Errors:  "payload/unprocessable-entity",
		})
	}

	errValidation := v.Validator.ValidateStruct(request)
	if errValidation != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Errors:  errValidation,
		})
	}

	exists, err := v.UseCase.Check(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Errors:  err.Error(),
		})
	}

	logger.Trace("END")

	return ctx.Status(fiber.StatusOK).JSON(model.CheckVoucherResponse{
		Success: true,
		Exists:  exists,
	})
}

// Generate creates vouchers with randomly selected seats for a flight
// @Summary Generate vouchers with random seat assignment
// @Description This endpoint generates vouchers for a specific flight with randomly chosen seats. It ensures fair distribution of seats and prevents manual seat selection bias. The system will automatically assign available seats to the generated vouchers.
// @Tags Vouchers
// @Accept json
// @Produce json
// @Param request body model.GenerateVoucherRequest true "Request body containing flight details and number of vouchers to generate"
// @Success 200 {object} model.GenerateVoucherResponse "Successfully generated vouchers with assigned seats"
// @Failure 400 {object} model.ErrorResponse "Bad request - invalid payload, unprocessable entity, validation error, or insufficient available seats"
// @Failure 500 {object} model.ErrorResponse "Internal server error"
// @Router /generate [post]
func (v *VoucherHandler) Generate(ctx *fiber.Ctx) error {
	method := "VoucherHandler.Generate"
	logger := v.Log.WithField("method", method)
	logger.Trace("BEGIN")

	request := new(model.GenerateVoucherRequest)
	if err := ctx.BodyParser(request); err != nil {
		logger.Error("failed parse body: ", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Errors:  "payload/unprocessable-entity",
		})
	}

	errValidation := v.Validator.ValidateStruct(request)
	if errValidation != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ErrorResponse{
			Success: false,
			Errors:  errValidation,
		})
	}

	seats, err := v.UseCase.Generate(ctx.UserContext(), request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"errors":  err.Error(),
		})
	}

	logger.Trace("END")

	return ctx.Status(fiber.StatusOK).JSON(model.GenerateVoucherResponse{
		Success: true,
		Seats:   seats,
	})
}
