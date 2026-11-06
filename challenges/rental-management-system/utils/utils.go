package utils

import (
	"github.com/atlas-devel/Go-Journey/challenges/rental-management-system/data"
	"github.com/atlas-devel/Go-Journey/challenges/rental-management-system/models"
)


func DisplayProperties(properties []models.Property)[]models.Property{ 
	return data.Properties

}

func FindProperty(id int, properties []models.Property)(models.Property, bool){
	for _,prop:=range properties{
		if prop.ID==id{
			return prop,true
		}
	}
	return models.Property{},false
}