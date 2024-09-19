package entity

import "time"

type CustomersTransaction struct {
	ContractNumber    string
	CustomerId        int
	OTR               float64
	AdminFee          float64
	InstallmentAmount float64
	InterestAmount    float64
	AssetName         string
	TransactionDate   time.Time
}
