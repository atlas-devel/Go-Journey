package main

import (
	"fmt"
	"strconv"

	"github.com/atlas-devel/Go-Journey/funcs/helpers"
)

func main() {

	total := helpers.MakeAdd(helpers.Add, 10, 12)
	fmt.Println(total)

	// a short int to string converter practice
	age := 20
	message := "He is " + strconv.Itoa(age)
	fmt.Println(message)

}
