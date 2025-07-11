package usecase

import (
	"bookcabin-backend/internal/entity"
	"bookcabin-backend/internal/model"
	"bookcabin-backend/internal/repository"
	"context"
	"fmt"
	"sort"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AircraftUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	VoucherRepository *repository.VoucherRepository
}

func NewAircraftUseCase(
	db *gorm.DB,
	log *logrus.Logger,
	voucherRepository *repository.VoucherRepository,
) *AircraftUseCase {
	return &AircraftUseCase{
		DB:                db,
		Log:               log,
		VoucherRepository: voucherRepository,
	}
}

func (a *AircraftUseCase) ListSeats(ctx context.Context, request *model.ListAircraftSeatsRequest) []model.AircraftSeatResponse {
	method := "AircraftUseCase.ListSeats"
	logger := a.Log.WithField("method", method)
	logger.Trace("BEGIN")
	logger.WithField("request", request).Debug("request-data")

	aircrafts := entity.NewAircraft()
	aircraftConfig := aircrafts.GetAircraftConfig(request.AircraftType)

	vouchers := make([]entity.Voucher, 0)
	err := a.VoucherRepository.FindByAircraftType(a.DB, &vouchers, request.AircraftType)
	if err != nil {
		panic(err)
	}

	logger.Info(fmt.Sprintf("found %d assigned seats", len(vouchers)*3)) // because 1 row db = 3 seats

	// Buat map kursi yang sudah dipakai
	assignedSeats := make(map[string]bool)
	for _, v := range vouchers {
		if v.Seat1 != "" {
			assignedSeats[v.Seat1] = true
		}
		if v.Seat2 != "" {
			assignedSeats[v.Seat2] = true
		}
		if v.Seat3 != "" {
			assignedSeats[v.Seat3] = true
		}
	}

	// Bangun response berdasarkan MaxRow dan Seats
	seats := make([]model.AircraftSeatResponse, 0)
	for row := 1; row <= aircraftConfig.MaxRow; row++ {
		for _, seat := range aircraftConfig.Seats {
			seatCode := fmt.Sprintf("%d%s", row, seat)
			seats = append(seats, model.AircraftSeatResponse{
				RowNumber: row,
				Seat:      seat,
				Assigned:  assignedSeats[seatCode],
			})
		}
	}

	sort.Slice(seats, func(i, j int) bool {
		if seats[i].RowNumber == seats[j].RowNumber {
			return seats[i].Seat < seats[j].Seat
		}
		return seats[i].RowNumber < seats[j].RowNumber
	})

	logger.Trace("END")
	return seats
}
