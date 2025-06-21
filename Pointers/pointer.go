package main

import "fmt"

func modify(a *int) {
	*a = *a * 2
}

func main() {
	// var a int
	a := 10
	fmt.Println(a)

	// var ptr *int
	ptr := &a
	// fmt.Println(ptr)
	// fmt.Println(*ptr)
	modify(ptr)
	fmt.Println(a)
}
