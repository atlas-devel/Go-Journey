package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := "123"
	y, err := strconv.Atoi(x)
	fmt.Println(y,err)

}