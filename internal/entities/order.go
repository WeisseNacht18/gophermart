package entities

type Order struct {
	ID      string  `json:"order"`
	Status  string  `json:"status"`
	Accrual float64 `json:"accrual"`
}
