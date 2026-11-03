package utils

import (
	"github.com/atlas-devel/Go-Journey/challenges/data"
)

func GetStudentById(id string) string {

	name, ok := data.Students[id]
	if !ok {
		return "Student not found"
	} else {
		return name
	}

}
