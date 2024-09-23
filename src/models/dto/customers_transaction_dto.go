package dto

type CreateCustomerTransactionDto struct {
	ContractNumber    string  `json:"contract_number"`
	CustomerId        int     `validate:"required" json:"customer_id"`
	Otr               float64 `validate:"required" json:"otr"`
	AdminFee          float64 `validate:"required" json:"admin_fee"`
	InstallmentAmount float64 `validate:"required" json:"installment_amount"`
	InterestAmount    float64 `validate:"required" json:"interest_amount"`
	AssetName         string  `validate:"required" json:"asset_name"`
}

type CustomerTransactionDto struct {
	ContractNumber    string  `json:"contract_number,omitempty"`
	CustomerId        int     `json:"customer_id,omitempty"`
	Otr               float64 `json:"otr,omitempty"`
	AdminFee          float64 `json:"admin_fee,omitempty"`
	InstallmentAmount float64 `json:"installment_amount,omitempty"`
	InterestAmount    float64 `json:"interest_amount,omitempty"`
	AssetName         string  `json:"asset_name,omitempty"`
}
