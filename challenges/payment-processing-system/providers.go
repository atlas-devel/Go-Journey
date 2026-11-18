package main

import (
	"fmt"
)

// MTN structs methods

func (m MTN) ProcessPayment(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Cannot process payment due to invalid amount: %f", amount)
	}
	ProcessingResponse("MTN", ResponsePurpose{Log: false, PaymentProcessing: true}, amount)
	return nil
}

func (m MTN) LogTransaction(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Cannot log transaction due to invalid amount: %f", amount)
	}
	ProcessingResponse("MTN", ResponsePurpose{Log: true, PaymentProcessing: false}, amount)
	return nil
}

func (m MTN) ValidatePayment(amount float64) (bool, error) {
	if amount <= 0 {
		return false, fmt.Errorf("Cannot validate payment due to invalid amount: %f", amount)
	}
	if amount > 500000{
		ProcessingResponse("MTN", ResponsePurpose{Log: false, PaymentProcessing: false}, amount)
		return false,fmt.Errorf("Amount exceed maximum limit of 500 000RWF: %f", amount)
	}
	return true, nil
}



// Airtel structs methods

func (m Airtel) ProcessPayment(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Cannot process payment due to invalid amount: %f", amount)
	}
	ProcessingResponse("Airtel", ResponsePurpose{Log: false, PaymentProcessing: true}, amount)
	return nil
}

func (m Airtel) LogTransaction(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Cannot log transaction due to invalid amount: %f", amount)
	}
	ProcessingResponse("Airtel", ResponsePurpose{Log: true, PaymentProcessing: false}, amount)
	return nil
}

func (m Airtel) ValidatePayment(amount float64) ( bool,error) {
		if amount <= 0 {
		return false, fmt.Errorf("Cannot validate payment due to invalid amount: %f", amount)
	}
	return true,nil
}



// BankTransfer structs methods
func (m BankTransfer) ProcessPayment(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Cannot process payment due to invalid amount: %f", amount)
	}
	ProcessingResponse("BankTransfer", ResponsePurpose{Log: false, PaymentProcessing: true}, amount)
	return nil
}

func (m BankTransfer) LogTransaction(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("Cannot log transaction due to invalid amount: %f", amount)
	}
	ProcessingResponse("BankTransfer", ResponsePurpose{Log: true, PaymentProcessing: false}, amount)
	return nil
}

func (m BankTransfer) ValidatePayment(amount float64) ( bool,error) {
		if amount <= 0 {
		return false, fmt.Errorf("Cannot validate payment due to invalid amount: %f", amount)
	}
	return true,nil
}


//interface function to process transaction

func ProcessTransaction(amount float64, service PaymentService)error{
 	isValid,err:=service.ValidatePayment(amount)
	
	if err!=nil{
		return err
	}
	if !isValid {
    return fmt.Errorf("payment validation failed")
}

	if err:=service.LogTransaction(amount);err!=nil{
		return err
	}
	if err:=service.ProcessPayment(amount);err!=nil{
		return err
	}

	return  nil
}
// test standalon log only

func LogOnly(amount float64, service TransactionLogger)error{
	if err:=service.LogTransaction(amount);err!=nil{
		return err
	}
	return nil
}