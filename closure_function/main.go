package main

import "fmt"

func main() {
	fn := newCounterFunc()
	fmt.Println(fn())
	fmt.Println(fn())
	fmt.Println(fn())
}

func newCounterFunc() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}