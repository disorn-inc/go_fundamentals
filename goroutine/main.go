package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(3)
	start := time.Now()
	go doSomething(wg)
	go doSomething(wg)
	go doSomething(wg)
	// time.Sleep(120 * time.Millisecond)
	wg.Wait()
	// fmt.Println(time.Since(start))
	// example()
	fmt.Println(time.Since(start))
}

func doSomething(wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("something")
	wg.Done()
}

func example() {
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			fmt.Print("-")
		}
	}()
	go func ()  {
		for {
			time.Sleep(200 * time.Millisecond)
			fmt.Print("+")
		}
	}()
	time.Sleep(time.Second)
}
