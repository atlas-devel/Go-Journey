package methods

import (
	"fmt"
	"strconv"
)

func StringToInt(s string) (int, error) {
	// using strconv package to convert string to int
	y, err := strconv.Atoi(s)
	fmt.Println(y,err)
	return y, err
}