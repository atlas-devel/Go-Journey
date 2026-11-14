package main

import "fmt"

func main() {
	score := 50
	scorePointer := &score
	fmt.Println(scorePointer)
	fmt.Println(*scorePointer)
	*scorePointer = 100
	fmt.Println(score)

}
