package main

import "fmt"

type ResponsePurpose struct{
Log bool
PaymentProcessing bool

}

func ProcessingResponse(serviceProvider string,purpose ResponsePurpose,amount float64) {
 if purpose.Log {
  fmt.Printf("%s Transaction logged: %.2f\n", serviceProvider, amount)
 }
 if purpose.PaymentProcessing {
  fmt.Printf("%s Payment processing: %.2f\n", serviceProvider, amount)
 }
}