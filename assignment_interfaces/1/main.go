package main

import "fmt"

type shape interface {
	calculateArea() float64
}
type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLenght float64
}

func main() {
	t := triangle{height: 3.2, base: 1.0}
	s := square{sideLenght: 2.42}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(s.calculateArea())
}

func (t triangle) calculateArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) calculateArea() float64 {
	return s.sideLenght * s.sideLenght
}
