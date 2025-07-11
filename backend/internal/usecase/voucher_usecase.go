package usecase

import (
	"bookcabin-backend/internal/entity"
	"bookcabin-backend/internal/model"
	"bookcabin-backend/internal/repository"
	"bookcabin-backend/internal/util"
	"context"
	"fmt"

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

	db := v.DB.WithContext(ctx)
	isExists, err := v.VoucherRepository.IsExists(db, request.FlightNumber, request.FlightDate)
	if err != nil {
		panic(err)
	}

	if isExists {
		return nil, fmt.Errorf("vouchers/already-exists")
	}

	assignedSeats, err := v.VoucherRepository.FindAssignedSeats(db, request.FlightDate, request.AircraftType)
	if err != nil {
		panic(err)
	}

	aircrafts := entity.NewAircraft()
	aircraftConfig := aircrafts.GetAircraftConfig(request.AircraftType)

	seats, err := util.GenerateUniqueSeats(assignedSeats, aircraftConfig)
	if err != nil {
		return nil, err
	}

	voucher := entity.NewVoucher(&entity.CreateVoucherProps{
		CrewID:       request.CrewID,
		CrewName:     request.CrewName,
		FlightNumber: request.FlightNumber,
		FlightDate:   request.FlightDate,
		AircraftType: request.AircraftType,
		Seat1:        seats[0],
		Seat2:        seats[1],
		Seat3:        seats[2],
	})

	if err = v.VoucherRepository.Create(db, voucher); err != nil {
		panic(err)
	}

	logger.Trace("END")
	return seats, nil
}
