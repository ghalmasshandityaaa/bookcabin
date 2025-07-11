package model

import "bookcabin-backend/internal/entity"

type ListAircraftSeatsRequest struct {
	AircraftType entity.AircraftType `json:"type" validate:"required,oneof=ATR 'Airbus 320' 'Boeing 737 Max'"`
}
