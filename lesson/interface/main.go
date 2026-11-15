package main

import "fmt"

type Player interface {
	Play()
}

type Phone struct {
	manufacturer string
	year         string
}

type Tv struct {
	manafacturer string
	year         string
}

type Speaker struct {
	manufacturer string
	year         string
}

func (p Phone) Play() {
	fmt.Println("Phone is beeping")
}
func (t Tv)Play(){
	fmt.Println("Tv is speaking")
}
func (s Speaker)Play(){
	fmt.Println("Speaker is speaking")
}


func startSpeaking(s Player){
	s.Play()
}

func main(){
	phone:=Phone{
		manufacturer: "Google pixel",
		year: "2026",
	}
	tv:=Tv{
		manafacturer: "Samsung",
		year: "2018",
	}
	speaker:=Speaker{
		manufacturer: "JBL",
		year: "2023",
	}
	startSpeaking(phone)
	startSpeaking(tv)
	startSpeaking(speaker)

}