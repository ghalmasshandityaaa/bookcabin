package model

type ErrorResponse struct {
	Success bool `json:"success"`
	Errors  any  `json:"errors,omitempty"`
}

type WebResponse[T any] struct {
	Ok     bool        `json:"ok"`
	Data   T           `json:"data,omitempty"`
	Errors interface{} `json:"errors,omitempty"`
}
