package main

import "fmt"

func addBonus(salary *int) {
	*salary +=20
}

func main() {
	salary := 120
	addBonus(&salary)
	fmt.Println(salary)
}