package main

import "fmt"

func main() {
	ch := make(chan int)
	// 2 variable type use to send signal it should not use memory ex. struct{} and [0]func
	qCh := make(chan struct{})

	go fibonacci(ch, qCh)

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	qCh <- struct{}{}
	<-qCh
	fmt.Println("end")
}

func fibonacci(ch chan int, qCh chan struct{}) {
	a, b := 0, 1

	for {
		select {
		case ch <- a:
			a, b = b, a+b
		case <- qCh:
			qCh <- struct{}{}
			return
		}
	}
}