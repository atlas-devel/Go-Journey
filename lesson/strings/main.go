package main

import "fmt"

func main() {
	sentence:="Hello there!!"
	for _, char:=range sentence{
		fmt.Printf("%v",string(char))
	}
}
