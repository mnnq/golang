package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {

	a := triangle{height: 2, base: 6}
	b := square{sideLength: 4.5654}

	printArea(b)
	printArea(a)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (d triangle) getArea() float64 {
	return 0.5 * d.base * d.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
