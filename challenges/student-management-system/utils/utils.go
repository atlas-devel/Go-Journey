package utils

import "github.com/atlas-devel/Go-Journey/challenges/sdms/models"


func CalculateAverage(student models.Student) float64 {

	totalStudentMarks := 0

	for _, mark := range student.Marks {
		totalStudentMarks = totalStudentMarks + int(mark)
	}

	return float64(totalStudentMarks / len(student.Marks))

}

func  HasPassed(student models.Student)bool{
	marks:=CalculateAverage(student)
	var decision bool

	if marks>=50{
		 decision=true
	}else{
		 decision=false
	}
	return decision
	

}

func UpdateStudentName(){
	
}