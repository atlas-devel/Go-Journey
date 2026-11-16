package main

import (
	"errors"
	"fmt"
	"strings"
)

type Sender interface {
	Send(message string)error
}
type Logger interface {
	Log(message string)error
}

type Validator interface {
	Validate(message string)(bool,error)
}

type NotificationService interface {
	Sender
	Logger
	Validator
}

// structs

type Email struct{}

func (e Email) Send(message string)error {
	if strings.TrimSpace(message)==""{
     return errors.New("Invalid message")
	}
	fmt.Println("Sending Email: ",message)
	return  nil
}


func (e Email) Log(message string)error {
	if strings.TrimSpace(message)==""{
     return errors.New("Invalid message")
	}
	fmt.Println("Logging Email: ",message)
	return  nil

}
func (e Email)Validate(message string)(bool,error){
	if strings.TrimSpace(message)==""{
     return false,errors.New("message can't be empty")
	}
	return  true,nil

}


// sms
type SMS struct{}

func (e SMS) Send(message string)error {
	if strings.TrimSpace(message)==""{
     return errors.New("Invalid message")
	}
	fmt.Println("Sending SMS: ",message)
	return  nil
}


func (e SMS) Log(message string)error {
	if strings.TrimSpace(message)==""{
     return errors.New("Invalid message")
	}
	fmt.Println("Logging SMS: ",message)
	return  nil

}
func (e SMS)Validate(message string)(bool,error){
	if strings.TrimSpace(message)==""{
		
     return false,errors.New("message can't be empty")
	}
	return  true,nil

}



//test single sender only without validate and logs

func TestMain(message string,s Sender){
	s.Send(message)

}