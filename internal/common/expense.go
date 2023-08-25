package common

type Expense struct {
	ID        string  `json:"id"`
	CreatedAt int64   `json:"created_at"`
	Amount    float64 `json:"amount"`
	Tax       float64 `json:"tax"`
}
