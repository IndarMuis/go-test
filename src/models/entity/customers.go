package entity

import "time"

type Customers struct {
	Id          int
	Nik         string
	FullName    string
	LegalName   string
	BirthPlace  string
	BirthDate   string
	Salary      string
	KTPPhoto    string
	SelfiePhoto string
	CreatedAt   time.Time
	IsSuspend   int
}
