package model

type ErrorResponse struct {
	Success bool `json:"success"`
	Errors  any  `json:"errors,omitempty"`
}
