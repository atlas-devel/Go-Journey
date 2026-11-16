package main

import (
	"fmt"
)

func ProcessNotification(message string, service NotificationService) error {
	valid, err := service.Validate(message)
	if err != nil {
		return err
	}
	if !valid {
		return err
	}

	if err:=service.Log(message);err!=nil{
		return err
	}
	if err:=service.Send(message);err!=nil{
		return err
	}

	return nil
}

func main() {
	if err:=ProcessNotification("Hello, World!", Email{});err!=nil{
		fmt.Println("Error:", err)
	}
	if err:=ProcessNotification("Leon testing!",SMS{});err!=nil{
		fmt.Println("Error:", err)
	}

	TestMain("Test sender!",Email{})
	
}