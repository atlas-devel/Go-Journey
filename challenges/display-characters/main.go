package main

import "fmt"

func main() {
str:="Hello World"

for index:=0;index<len(str);index++{
	fmt.Println(string(str[index]))
}

}
