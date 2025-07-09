package repository

import (
	"bookcabin-backend/internal/entity"

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
