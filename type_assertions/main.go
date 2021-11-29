package main

import "fmt"

func main() {
	var a, b Obj
	a = rectangle{w: 4, l: 4}
	b = triangle{w: 4, h: 4}

	fmt.Println(a.Area())
	fmt.Println(b.Area())

	if v,ok := b.(triangle); ok {
		v.h = 5
	}
}

type Obj interface {
	Area() float64 
}

type rectangle struct {
	w, l float64
}

func (rec rectangle) Area() float64 {
	return rec.l * rec.w
}

type triangle struct {
	w, h float64
}

func (tri triangle) Area() float64 {
	return 0.5 * tri.h * tri.w
}
