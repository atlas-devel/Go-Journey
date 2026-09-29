package preview

import "fmt"

func Preview() {
	arrayStrings := [...]string{"Leon", "Amina", "Alia", "Sonia"}
	
	for _, item:=range arrayStrings{
fmt.Println(item)
	}

}
