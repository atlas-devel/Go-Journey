package main

import "fmt"

func main() {

	// using var
	var a float32 = 127.3
	var name  = "John"
	var age int = 30
	var isStudent bool = true

	// use const 
	const Score=20

	// use implicity variables
	email:="irakaramale@gmail.com"
	
	fmt.Printf("%s\n",name)//%s: print string value
	fmt.Printf("%d\n",age)//%d: print integer value
	fmt.Printf("%t\n",isStudent)//%t: print type of the specified variable
	fmt.Printf("%f\n",a)//%f: print float value
	fmt.Printf("Email: %v", email)//%s: print defaul representation
	
}