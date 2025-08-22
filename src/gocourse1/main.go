package main

import (
	"fmt"
	helper "gocourse1/utils"
	"time"
)

var myArray [5]int
var myArray2 = [4]int{1, 2, 3, 4}

func main() {

	fmt.Println("Hello, GOphers!")
	gophers := helper.Gophers
	helper.Printer("Hello" + ", " + gophers)
	myArray[1] = 4

	for _, v := range myArray2 {
		fmt.Println(v)
	}

	fmt.Println("Summ of 2 and 4 is ", summ(2, 4))

	fmt.Println("More complexed tasks on the Go:")
	testGoRoutine()

}

func testGoRoutine() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)
}
