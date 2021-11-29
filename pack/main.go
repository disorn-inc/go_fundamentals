package main

import (
	"fmt"
	"github/disorn-inc/go_fundamentals/pack/book"
)

func main() {
	b := book.New()
	fmt.Printf("%T %v\n", b, b)

	b.Name = "Ring"
	fmt.Printf("%T %v\n", b, b)
}