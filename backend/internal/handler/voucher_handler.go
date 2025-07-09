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
