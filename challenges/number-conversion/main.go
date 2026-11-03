package main

import (
	"fmt"

	"github.com/atlas-devel/Go-Journey/challenges/number-conversion/utils"
)

func main() {
	res, err := utils.ConvertNumber("hello")
	if err != nil {
		fmt.Println("Error: ",err)
	} else {
		fmt.Println(res)
	}
}
