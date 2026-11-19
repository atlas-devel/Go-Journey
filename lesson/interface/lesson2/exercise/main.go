package main

import "fmt"

func IdentifyValue(value any) string {
	if value, ok := value.(string); ok {
		return fmt.Sprintf("String: %s",value)
	}
	if value, ok := value.(int); ok {
		return fmt.Sprintf("Integer: %d",value)
	}
	if value, ok := value.(bool); ok {
		return fmt.Sprintf("Boolean: %t",value)
	}
	return "Unknown Type"
}

func main() {
	fmt.Println(IdentifyValue("Leon"))
	fmt.Println(IdentifyValue(22))
	fmt.Println(IdentifyValue(true))
	fmt.Println(IdentifyValue(12.5))

}