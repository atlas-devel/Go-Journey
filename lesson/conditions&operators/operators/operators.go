package operators

import "fmt"

func Operators() {
	x := uint(8)
	y := uint(11)
	z:=2>x
	var result=x==y
		// fmt.Printf("Type of x is %T\n", x)
		// fmt.Println(y)
		fmt.Printf("Result of x==y is %T\n", result)
		fmt.Printf("Result of 2>x is %v\n", z)

}