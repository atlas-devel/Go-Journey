package switchstatement

import "fmt"

func SwitchStatement() {
	days:=""

	switch days{
	case "Monday":
		fmt.Println("kuwa 1")
	case "Tuesday":
		fmt.Println("kuwa 2")
	case "Wednesday":
		fmt.Println("kuwa 3")
	case "Thursday":
		fmt.Println("kuwa 4")
	case "Friday":
		fmt.Println("kuwa 5")
	case "Saturday":
		fmt.Println("kuwa 6")
	case "Sunday":
		fmt.Println("kucyumweru")
	default:
		fmt.Println("Please choose any day of the week")
	}
	
	
}