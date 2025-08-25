package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Goroutines: ")
	testGoRoutine()
	fmt.Println()

	fmt.Println("Buffered channels: ")
	testGoBufChan()
	fmt.Println()

	fmt.Println("Sync Wait group usage: ")
	testWaitGroup()
	fmt.Println()

	fmt.Println("Channels with future: ")
	testGoChan()
	fmt.Println()
}

func testGoRoutine() {
	// так в Го не делают
	// var ch chan string
	// ch = make(chan string)
	// нужно вот так, иначе линтер ругается
	ch := make(chan string)
	fmt.Printf("Value of chan is %T\n", ch)
	fmt.Printf("Address of chan is %v\n", ch)
	go worker(ch)
	time.Sleep(10 * time.Millisecond)

	msg := <-ch
	fmt.Println("Message from Goroutine: ", msg)
}

func worker(ch chan string) {
	ch <- "Работа завершена"
}

func greeter(c chan string) {
	fmt.Println("Hello ", <-c, "!")
}

func testGoChan() {

	ch := make(chan string)

	go greeter(ch)

	ch <- "Timmy"

	fmt.Println("End of func")
}

func testGoBufChan() {
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "gophers"
	// чтобы в цикле обработать содержимое канала его нужно
	// сначала закрыть
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
	// fmt.Println(<-ch)
	// fmt.Println(<-ch)
}

func testWaitGroup() {

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go task(i)

	}
	wg.Wait()
}

func task(id int) {
	defer wg.Done()
	fmt.Printf("Задача %d выполнена сполна\n", id)
}
