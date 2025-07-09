package usecase

import (
	"bookcabin-backend/internal/model"
	"bookcabin-backend/internal/repository"
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type VoucherUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	VoucherRepository *repository.VoucherRepository
}

func NewVoucherUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	voucherRepository *repository.VoucherRepository,
) *VoucherUseCase {
	return &VoucherUseCase{
		DB:                db,
		Log:               log,
		VoucherRepository: voucherRepository,
	}
}

func (v *VoucherUseCase) Check(ctx context.Context, request *model.CheckVoucherRequest) (bool, error) {
	method := "VoucherUseCase.Check"
	logger := v.Log.WithField("method", method)
	logger.Trace("BEGIN")

	db := v.DB.WithContext(ctx)
	isExists, err := v.VoucherRepository.IsExists(db, request.FlightNumber, request.FlightDate)
	if err != nil {
		panic(err)
	}

	logger.Trace("END")
	return isExists, nil
}

func (v *VoucherUseCase) Generate(ctx context.Context, request *model.GenerateVoucherRequest) ([]string, error) {
	method := "VoucherUseCase.Generate"
	logger := v.Log.WithField("method", method)
	logger.Trace("BEGIN")

	logger.Trace("END")
	return []string{}, nil
}
