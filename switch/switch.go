package main

import "fmt"

func main() {
	// day := 2
	// switch day {
	// case 1:
	// 	fmt.Println("1")

	// case 2:
	// 	fmt.Println("2")

	// case 3:
	// 	fmt.Println("3")

	// default:
	// 	fmt.Println("Others")
	// }

	month := "January"

	switch month {
	case "January", "February", "March":
		fmt.Println("Good")

	default:
		fmt.Println("Bad")
	}
}
