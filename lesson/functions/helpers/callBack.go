package helpers

import "fmt"

func MakeAdd(callback func(uint,uint) uint,num1 uint, num2 uint) string {
	total := callback(num1,num2)
	return fmt.Sprintf("The total is %d",total)
}