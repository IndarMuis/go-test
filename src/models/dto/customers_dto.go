package dto

import "time"

type CustomerDto struct {
	Nik         string `json:"nik,omitempty"`
	FullName    string `json:"full_name,omitempty"`
	LegalName   string `json:"legal_name,omitempty"`
	BirthPlace  string `json:"birth_place,omitempty"`
	BirthDate   string `json:"birth_date,omitempty"`
	Salary      string `json:"salary,omitempty"`
	KTPPhoto    string `json:"ktp_photo,omitempty"`
	SelfiePhoto string `json:"selfie_photo,omitempty"`
}

type CreateCustomerDto struct {
	Nik         string `validate:"required,len=16,number" json:"nik,omitempty"`
	FullName    string `validate:"required" json:"full_name,omitempty"`
	LegalName   string `validate:"required" json:"legal_name,omitempty"`
	BirthPlace  string `validate:"required" json:"birth_place,omitempty"`
	BirthDate   string `validate:"required" json:"birth_date,omitempty"`
	Salary      string `validate:"required,number" json:"salary,omitempty"`
	KTPPhoto    string `validate:"required,url" json:"ktp_photo,omitempty"`
	SelfiePhoto string `validate:"required,url" json:"selfie_photo,omitempty"`
}

type UpdateCustomerDto struct {
	FullName    string `json:"full_name,omitempty"`
	LegalName   string `json:"legal_name,omitempty"`
	BirthPlace  string `json:"birth_place,omitempty"`
	BirthDate   string `json:"birth_date,omitempty"`
	Salary      string `validate:"number" json:"salary,omitempty"`
	KTPPhoto    string `validate:"url" json:"ktp_photo,omitempty"`
	SelfiePhoto string `validate:"url" json:"selfie_photo,omitempty"`
	IsSuspend   int    `validate:"url" json:"is_suspend,omitempty"`
}

type CustomerDetailsDto struct {
	Nik         string    `json:"nik,omitempty"`
	FullName    string    `json:"full_name,omitempty"`
	BirthPlace  string    `json:"birth_place,omitempty"`
	BirthDate   string    `json:"birth_date,omitempty"`
	Salary      string    `json:"salary,omitempty"`
	KTPPhoto    string    `json:"ktp_photo,omitempty"`
	SelfiePhoto string    `json:"selfie_photo,omitempty"`
	CustomerId  int       `json:"customer_id,omitempty"`
	LegalName   string    `json:"legal_name,omitempty"`
	TenorLimit  float64   `json:"tenor-limit,omitempty"`
	StartDate   string    `json:"start_date,omitempty"`
	EndDate     string    `json:"end-date,omitempty"`
	AmountUsed  float64   `json:"amount-used,omitempty"`
	IsSuspend   int       `json:"is_suspend,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
