package main

import "fmt"

func main() {
	vehicles := []string{
		"car", "bus", "car", "motorcycle",
		"bus", "car", "truck", "motorcycle",
	}

	sortedVehicles := make(map[string]int)
	var totalVehicles int = 0

	// sort vehicles
	for _, value := range vehicles {
		_, exists := sortedVehicles[value]
		if exists {
			sortedVehicles[value]++
		} else {
			sortedVehicles[value] = 1
		}
	}

	fmt.Println("----------------VEHICLES COUNT---------------------------")
	fmt.Println(sortedVehicles)
	fmt.Println("----------------EACH VEHICLE TYPE---------------------------")

	// each vehicle type and count and their total
	for key, value := range sortedVehicles {
		fmt.Print("Vehicle type: ", key, ", Count: ", value,"\n")
		totalVehicles = totalVehicles + value
	}

	fmt.Println("Total vehicles: ",totalVehicles)

}
