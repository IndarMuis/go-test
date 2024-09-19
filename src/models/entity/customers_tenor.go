package entity

type CustomersTenor struct {
	Id         int
	CustomerId int
	TotalMonth int
	TenorLimit float64
	StartDate  string
	EndDate    string
	Status     string
}
