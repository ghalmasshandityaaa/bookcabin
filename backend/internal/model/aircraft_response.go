package model

type AircraftSeatResponse struct {
	RowNumber int    `json:"row_number"`
	Seat      string `json:"seat"`
	Assigned  bool   `json:"assigned"`
}

type AircraftSeatSwaggerResponse struct {
	Ok   bool                   `json:"ok"`
	Data []AircraftSeatResponse `json:"data"`
}
