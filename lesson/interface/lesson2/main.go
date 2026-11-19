package main

import "fmt"

type MyStringer interface {
	makeString() string
}

type Person struct{
	Name string
	Email string
}

func (p Person) makeString() string {
	return fmt.Sprintf("Person - Name: %s, Person - Email: %s", p.Name, p.Email)
}

func Display(value any){
	if v, ok:=value.(MyStringer); ok{
		fmt.Println(ok)
		fmt.Println(v.makeString())
	}else{
		fmt.Println("not a MyStringer:", value)
	}	
}

func main() {
 Display("Hello World")
 Display(10)
 Display(3.14)
 Display(true)
 Display(Person{Name: "John Doe", Email: "john.doe@example.com"})

}
