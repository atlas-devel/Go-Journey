package preview

import "fmt"

func Preview() {
	arrayStrings := [...]string{"Leon", "Amina", "Alia", "Sonia"}

	sl := arrayStrings[2:]
	sl[0] = "Sangano"
	fmt.Println(arrayStrings, sl)

	// 	for _, item:=range arrayStrings{
	// fmt.Println(item)
	// 	}

}
