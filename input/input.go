package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var name string
	// fmt.Print("Enter your name: ")
	// fmt.Scan(&name)
	// fmt.Printf("Hello %s ", name)

	reader := bufio.NewReader(os.Stdin)
	name1, _ := reader.ReadString('\n')
	fmt.Println(name1)

}
