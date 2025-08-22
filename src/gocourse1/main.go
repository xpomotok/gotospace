package main

import (
	"fmt"
	helper "gocourse1/utils"
)

func main() {

	fmt.Println("Hello, GOphers!")
	gophers := helper.Gophers
	helper.Printer("Hello" + ", " + gophers)
}
