package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	counter := 0
	for {
		fmt.Println(counter)
		counter++
		if counter == 3 {
			break
		}
	}
}
