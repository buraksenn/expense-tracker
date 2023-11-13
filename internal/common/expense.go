package common

type Expense struct {
	ID        string  `json:"id" dynamodbav:"id"`
	CreatedAt int64   `json:"created_at" dynamodbav:"created_at"`
	Amount    float64 `json:"amount" dynamodbav:"amount"`
	Tax       float64 `json:"tax" dynamodbav:"tax"`
}
