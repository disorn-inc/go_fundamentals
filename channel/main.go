package main

import "fmt"

func main() {
	ch := make(chan int)
	go fibonacci(ch)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func fibonacci(ch chan int) {
	a, b := 0, 1

	for {
		ch <- a
		a, b = b, a+b
	}
}