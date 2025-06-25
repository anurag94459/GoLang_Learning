package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error found!!", err)
		return
	}
	defer res.Body.Close()
	// fmt.Print("%T\n", res)
	fmt.Println(res)

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Error in reading")
		return
	}
	fmt.Println(data)
	
}
