package transaction

import "time"

//Transaction documentation
type Transaction struct {
	Title     string    `json:"title"`
	Amount    float32   `json:"amount"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
}

//Transactions documentation
type Transactions []Transaction
