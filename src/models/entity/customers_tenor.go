package entity

type CustomersTenor struct {
	Id               int
	CustomerId       int
	TotalMonth       int
	TenorAmountLimit float64
	AmountUsed       float64
	StartDate        string
	EndDate          string
	Status           string
}
