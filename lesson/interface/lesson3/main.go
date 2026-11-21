package main

import "fmt"

type Payment interface {
    Pay()
}

type MTN struct{}

func (m *MTN) Pay() {
    fmt.Println("Paying with MTN")
}

func main (){
    var mtn *MTN = nil
    var payment Payment = mtn
    fmt.Println(payment==mtn)
}

func GetMTNProvider() *MTN {
    return &MTN{}
}