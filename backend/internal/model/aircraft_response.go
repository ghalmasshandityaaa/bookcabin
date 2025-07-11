package model

type AircraftSeatResponse struct {
	RowNumber int    `json:"row_number"`
	Seat      string `json:"seat"`
	Assigned  bool   `json:"assigned"`
}
