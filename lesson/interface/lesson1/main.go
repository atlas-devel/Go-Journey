package main

import "fmt"

type Email struct{}

type SMS struct{}

type WhatsApp struct{}

func (e Email) send(message string) {
	fmt.Printf("sending email: %s",message)
}
func (s SMS) send(message string) {
	fmt.Printf("sending SMS: %s",message)

}
func (w WhatsApp) send(message string) {
	fmt.Printf("sending whatsapp message: %s",message)

}

func sendNotificatin(n Notifier,message string){
	n.send(message)
}

func main() {
	email:=Email{}
	sms:=SMS{}
	whatsapp:=WhatsApp{}

	sendNotificatin(email,"Hello Leon!")
	sendNotificatin(sms,"Hello Leon!")
	sendNotificatin(whatsapp,"Hello Leon!")
}