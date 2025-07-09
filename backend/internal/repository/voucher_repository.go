package repository

import (
	"bookcabin-backend/internal/entity"
	"bookcabin-backend/internal/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type VoucherRepository struct {
	Repository[entity.Voucher]
	Log *logrus.Logger
}

func NewVoucherRepository(log *logrus.Logger) *VoucherRepository {
	return &VoucherRepository{
		Log: log,
	}
}

func (v *VoucherRepository) IsExists(db *gorm.DB, flightNumber, flightDate string) (bool, error) {
	method := "VoucherRepository.IsExists"
	logger := v.Log.WithField("method", method)
	logger.Trace("BEGIN")

	var result int
	err := db.Debug().Raw("SELECT 1 FROM vouchers WHERE flight_number = ? AND flight_date = ? LIMIT 1", flightNumber, flightDate).Scan(&result).Error

	if err != nil {
		return false, err
	}

	logger.Trace("END")
	return result == 1, nil
}

func (v *VoucherRepository) FindAssignedSeats(db *gorm.DB, flightDate string, aircraftType model.AircraftType) ([]string, error) {
	method := "VoucherRepository.FindAssignedSeats"
	logger := v.Log.WithField("method", method)
	logger.Trace("BEGIN")

	var seats []string
	query := `
		SELECT seat1 AS seat FROM vouchers WHERE flight_date = ? AND aircraft_type = ?
		UNION
		SELECT seat2 FROM vouchers WHERE flight_date = ? AND aircraft_type = ?
		UNION
		SELECT seat3 FROM vouchers WHERE flight_date = ? AND aircraft_type = ?
	`

	err := db.Raw(query,
		flightDate, aircraftType,
		flightDate, aircraftType,
		flightDate, aircraftType,
	).Scan(&seats).Error

	if err != nil {
		logger.WithError(err).Error("failed to fetch unique seats")
		return nil, err
	}

	logger.Trace("END")
	return seats, nil
}
