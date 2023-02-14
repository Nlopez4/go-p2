package main

import (
	"fmt"

	"example.com/packages/myapp"
)


func main() {

	greeting := fmt.Sprintf("hello %s", "from main")
	
	fmt.Println(greeting)

	myapp.PrintHello()

}