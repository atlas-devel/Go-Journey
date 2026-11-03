package utils

import "strconv"

func ConvertNumber(value string) (int, error){
	num,err:=strconv.Atoi(value)
	return  num, err

}