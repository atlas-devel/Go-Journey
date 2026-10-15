package main

import "fmt"

type Person struct {
	name  string
	email string
	age   uint
}

func main() {

	p1 := Person{
		name: "Leon",
		age: 22,
		email: "irakaramale@gmail.com",
	}
	fmt.Println(p1)

}