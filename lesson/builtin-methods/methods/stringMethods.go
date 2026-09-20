package methods

import (
	"fmt"
	"strings"
)

func StringMethods(name,surname string) {

		// using String package 
	name=strings.ToUpper(name)// convert string to upper 
	// fmt.Println(name, )

	var letters=[]byte(surname)// convert string to byte slice	
	letters[0]=strings.ToUpper(string(letters[0]))[0]
	fmt.Println(string(letters))

}