package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Cannot create file")
	}
	defer file.Close()

	content := "Hello World"

	_, errs := io.WriteString(file, content)
	if errs != nil {
		fmt.Println("Cannot write in a file")
		return
	}
	fmt.Println("File Created!!")
}
