package dto

import "time"

type Customers struct {
	Nik         string    `json:"nik,omitempty"`
	FullName    string    `json:"full_name,omitempty"`
	LegalName   string    `json:"legal_name,omitempty"`
	BirthPlace  string    `json:"birth_place,omitempty"`
	BirthDate   time.Time `json:"birth_date,omitempty"`
	Salary      float64   `json:"salary,omitempty"`
	KTPPhoto    string    `json:"ktp_photo,omitempty"`
	SelfiePhoto string    `json:"selfie_photo,omitempty"`
	IsSuspend   int       `json:"is_suspend,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

type CustomerDetails struct {
	CustomerId int       `json:"customer_id,omitempty"`
	LegalName  string    `json:"legal_name,omitempty"`
	TenorLimit float64   `json:"tenor-limit,omitempty"`
	StartDate  string    `json:"start_date,omitempty"`
	EndDate    string    `json:"end-date,omitempty"`
	AmountUsed float64   `json:"amount-used,omitempty"`
	IsSuspend  int       `json:"is_suspend,omitempty"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
}
