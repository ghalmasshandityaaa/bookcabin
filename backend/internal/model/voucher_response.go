package model

type CheckVoucherResponse struct {
	Success bool `json:"success"`
	Exists  bool `json:"exists"`
}

type GenerateVoucherResponse struct {
	Success bool     `json:"success"`
	Seats   []string `json:"seats"`
}
