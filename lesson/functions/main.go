package main

import "fmt"

func add(num1 uint, num2 uint) uint {
	return num1 + num2

}

func main() {

	total := add(20, 3)
	fmt.Println(total)

}
