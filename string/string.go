package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "apple,banana,orange"
	split := strings.Split(a, ",")
	fmt.Println(split)

	b := "     apple             "
	space := strings.TrimSpace(b)
	fmt.Println(space)

	c := "applebananaorange"
	count1 := strings.Count(c, "apple")
	fmt.Println(count1)

	d := "hello"
	e := "world"
	f := strings.Join([]string{d, ",", e}, "")
	fmt.Println((f))
}
