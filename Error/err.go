package main

import "fmt"

func div(a float32, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("Not divisible by 0")
	}
	return a / b, nil
}

func main() {
	res, _ := div(3, 0)
	fmt.Println(res)
}
