package main

import "fmt"

func add(a int, b int) (res int) {
	// return a + b
	res = a + b
	return
}

func main() {

	var res1 = add(3, 4)
	fmt.Println(res1)
}
