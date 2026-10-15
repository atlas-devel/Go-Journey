package main

import "fmt"

type Student struct {
	Name  string
	Age   uint
	Grade string
}

func main() {

	st1 := Student{
		Name:  "Eric",
		Age:   19,
		Grade: "A",
	}

	fmt.Printf("Student name is %s\n",st1.Name)
	fmt.Printf("%s is %d old\n",st1.Name,st1.Age)
	fmt.Printf("he/she is in Grade  %s\n",st1.Grade)

}