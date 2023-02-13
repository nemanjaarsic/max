package model

type Transaction struct {
	ID            string
	UserID        string
	OperationType OperationType
	Amount        float32
	Timestamp     string
}

type OperationType int

const (
	Deposit OperationType = iota
	Withdraw
)
