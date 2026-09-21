package conditions

import "fmt"

func Conditions(){
	x:=0
	if x<2 {
		fmt.Println("x is less than 2")
	}else if x==3 {
		fmt.Println("x equals to 3")
	}else if x>5 {
		fmt.Println("x is far beyond reach")
	}else{
		fmt.Println("no answer found")
	}
	
}