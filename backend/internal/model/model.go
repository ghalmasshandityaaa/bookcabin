package model

type ErrorResponse struct {
	Success bool `json:"success"`
	Errors  any  `json:"errors,omitempty"`
}

type WebResponse[T any] struct {
	Success bool        `json:"success"`
	Data    T           `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}
