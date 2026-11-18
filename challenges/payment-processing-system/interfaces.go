package main

// interfaces

type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

type TransactionLogger interface {
	LogTransaction(amount float64) error
}

type PaymentValidator interface {
	ValidatePayment(amount float64) (bool, error)
}

type PaymentService interface {
	PaymentProcessor
	TransactionLogger
	PaymentValidator
}

// structs
type BankTransfer struct{}
type MTN struct{}
type Airtel struct{}
