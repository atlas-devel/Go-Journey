package preview

import "fmt"

func Preview() {
	arrayStrings := [...]string{"Leon", "Amina", "Alia", "Sonia","Anaise"}

	sl := arrayStrings[2:3]
	sl[0] = "Sangano"
	sl=sl[:3]
	// fmt.Println(arrayStrings, sl)

	// 	for _, item:=range arrayStrings{
	// fmt.Println(item)
	// 	}
	fmt.Println(sl,len(sl),cap(sl))
}
