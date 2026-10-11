package main

import "fmt"

func main() {

	vehicleTypes := map[string]int{
		"toyota": 100,
		"suzuki": 5,
		"ford":   10,
		"hyndai": 13,
		"kia":    90,
	}

	for key, value:=range vehicleTypes{
		fmt.Println(key,value)
	}

	// fmt.Println(vehicleTypes)

}
