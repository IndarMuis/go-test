package entity

import "time"

type Customers struct {
	Id          int
	Nik         string
	FullName    string
	LegalName   string
	BirthPlace  string
	BirthDate   time.Time
	Salary      float64
	KTPPhoto    string
	SelfiePhoto string
	CreatedAt   time.Time
	IsSuspend   int
}
