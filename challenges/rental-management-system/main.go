package main

import (
	"github.com/atlas-devel/Go-Journey/challenges/rental-management-system/data"
	"github.com/atlas-devel/Go-Journey/challenges/rental-management-system/utils"
)

func main() {
	allProperties := utils.DisplayProperties(data.Properties)
	foundProperty, isfound := utils.FindProperty(2, allProperties)

	if isfound {
		println("Property Found: ", foundProperty.Title)
	} else {
		println("Property not found")
	}

}
