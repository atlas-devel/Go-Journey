package main

import "fmt"

type Student struct {
	Name  string
	Age   uint
	Grade string
}


func (s Student)getStudentDataByName(name string)Student{
	dummyData:=[]Student{
	{
		Name:  "Eric",
		Age:   19,
		Grade: "A",
	},
	{
		Name:  "Amina",
		Age:   22,
		Grade: "B",
	},{
		Name:  "Assa",
		Age:   21,
		Grade: "C",
	},
}
	var result Student

	for _,value:=range dummyData{
		if value.Name==name{
			result=value
		}
	}
return result

}

func main() {



	st1 := Student{}
	
	personData:=st1.getStudentDataByName("Amina")
	fmt.Println(personData)


}