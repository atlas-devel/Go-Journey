package main

import "fmt"

type Email struct{}

type SMS struct{}

type WhatsApp struct{}

func (e Email) Send(message string) {
	fmt.Printf("sending email: %s\n",message)
}
func (s SMS) Send(message string) {
	fmt.Printf("sending SMS: %s\n",message)

}
func (w WhatsApp) Send(message string) {
	fmt.Printf("sending whatsapp message: %s\n",message)

}

func sendNotificatin(n Notifier,message string){
	n.Send(message)
}

func NewEmailNotifier() Email{return  Email{}}
func NewSMSNotifier() SMS{return SMS{}}
func NewWhatsAppNotifier() WhatsApp{return WhatsApp{}}

func main() {
	// email:=Email{}
	// sms:=SMS{}
	// whatsapp:=WhatsApp{}

	// sendNotificatin(email,"Hello Leon!")
	// sendNotificatin(sms,"Hello Leon!")
	// sendNotificatin(whatsapp,"Hello Leon!")


	sendNotificatin(NewWhatsAppNotifier(),"Hello Leon")
	sendNotificatin(NewSMSNotifier(),"Your account is ready")
	sendNotificatin(NewWhatsAppNotifier(),"Welcome to Go!")
}