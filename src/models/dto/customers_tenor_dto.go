package dto

type CreateCustomerTenorDto struct {
	CustomerId       int     `validate:"required,number" json:"customer_id"`
	TotalMonth       int     `validate:"required,min=1,max=12,number" json:"total_month"`
	TenorAmountLimit float64 `validate:"required,number" json:"tenor_amount_limit"`
	StartDate        string  `validate:"required" json:"start_date"`
	EndDate          string  `validate:"required" json:"end_date"`
}

type UpdateCustomerTenorDto struct {
	CustomerId       int     `validate:"required,number" json:"customer_id"`
	TenorAmountLimit float64 `validate:"required,number" json:"tenor_amount_limit"`
}

type ActivateTenorDto struct {
	CustomerId int    `validate:"required,number" json:"customer_id"`
	TotalMonth int    `validate:"required,min=1,max=12,number" json:"total_month"`
	Status     string `validate:"required" json:"status"`
}

type CustomerTenorDto struct {
	CustomerId       int     `json:"customer_id,omitempty"`
	TotalMonth       int     `json:"total_month,omitempty"`
	TenorAmountLimit float64 `json:"tenor_amount_limit,omitempty"`
	AmountUsed       float64 `json:"amount_used,omitempty"`
	StartDate        string  `json:"start_date,omitempty"`
	EndDate          string  `json:"end_date,omitempty"`
	Status           string  `json:"status,omitempty"`
}

type FindTenorByCustomerIdDto struct {
	CustomerId int `validate:"required,number" json:"customer_id"`
}
