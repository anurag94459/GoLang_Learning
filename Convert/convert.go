package main

import "fmt"

func main() {
	var a int
	a = 10

	var num float64 = float64(a)

	fmt.Println(num)
	fmt.Printf("%T", num)
}
