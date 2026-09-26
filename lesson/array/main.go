package main

import "fmt"

func main() {

	arrayStrings := [...][5]string{
		{
			"leon",
			"herve",
			"sefu",
			"eric",
			"malik",
		},
		{
			"anna",
			"sandra",
			"masama",
			"cyubahiro",
			"anaise",
		},
		{
			"kami",
			"kamanda",
			"constantine",
			"izere",
			"christine",
		},
		{
			"clemment",
			"grace",
			"leon",
			"maria",
			"penelope",
		},
		{
			"amina",
			"ineza",
			"uwamahoro",
			"isange",
			"elissa",
		},
	}

	// fmt.Printf("%T",arrayStrings)

	for _, char := range arrayStrings {
		fmt.Println(char)
		for _, item := range char {
			fmt.Println(item)
		}
	}
}
