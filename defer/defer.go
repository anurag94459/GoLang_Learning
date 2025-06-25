package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func main() {
	fmt.Println("Starting")
	c := add(3, 4)
	defer fmt.Println(c) //Incase of Multiple Defer it works in LIFO order
	fmt.Println("Middle")
	fmt.Println("Ending")
}
