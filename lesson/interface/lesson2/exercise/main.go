package main

import "fmt"

type Person struct {
	Name string
	Email string
}


func IdentifyValue(value any) string {
	switch v:=value.(type){
	case int:
		return fmt.Sprintf("Integer: %d",v)
	case string:
		return fmt.Sprintf("String: %s",v)
	case bool:
		return fmt.Sprintf("Boolean: %t",v)
	case float64:
		return fmt.Sprintf("Float: %.2f",v)
	case []string:
		return  fmt.Sprintf("Slice: %v",v)
	case Person:
		return  fmt.Sprintf("Person: %v",v)
	default:
		return fmt.Sprintf("Unknown type: %v",v)
	}

}

func main() {
	fmt.Println(IdentifyValue("Hello World"))                
	fmt.Println(IdentifyValue(22))                      
	fmt.Println(IdentifyValue(12.5))                    
	fmt.Println(IdentifyValue(true))                   
	fmt.Println(IdentifyValue([]string{"Go", "PHP"}))   
	fmt.Println(IdentifyValue(Person{Name: "Alice", Email: "alice@example.com"})) 
}