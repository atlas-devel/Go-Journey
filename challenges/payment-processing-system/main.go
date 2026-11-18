package main

import "fmt"

func main() {
	if err := ProcessTransaction(0, MTN{}); err != nil {
		fmt.Println("Error occurred:", err)
	}else{
		fmt.Println("MTN Transaction succeed!")
	}
	if err:=ProcessTransaction(-2,Airtel{});err!=nil{
		fmt.Println("Error occurred:", err)
	}else{
		fmt.Println("Airtel Transaction succeed!")
	}
	if err:=ProcessTransaction(200,BankTransfer{});err!=nil{
		fmt.Println("Error occurred:", err)
	}else{
		fmt.Println("Bank Transfer Transaction succeed!")
	}

	if err:=LogOnly(0,MTN{});err!=nil{
		fmt.Println("Error occurred:", err)
	}else{
		fmt.Println("Amount Logged successfull!")
	}
}