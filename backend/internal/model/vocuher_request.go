package model

type AircraftType string

const (
	ATR          AircraftType = "ATR"
	Airbus320    AircraftType = "Airbus 320"
	Boeing737Max AircraftType = "Boeing 737 Max"
)

type CheckVoucherRequest struct {
	FlightNumber string `json:"flightNumber" validate:"required,min=3,max=100"`
	FlightDate   string `json:"date" validate:"required,is-valid-date"`
}

type GenerateVoucherRequest struct {
	CrewID       string       `json:"id" validate:"required,numeric"`
	CrewName     string       `json:"name" validate:"required,alpha-with-space,min=2,max=100"`
	FlightNumber string       `json:"flightNumber" validate:"required,alphanum,min=3,max=100"`
	FlightDate   string       `json:"date" validate:"required,is-valid-date"`
	AircraftType AircraftType `json:"aircraft" validate:"required,oneof=ATR 'Airbus 320' 'Boeing 737 Max'"`
}
