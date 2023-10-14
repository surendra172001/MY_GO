package main

import "fmt"

type shape interface {
	area() float64
}

type square struct {
	sideLength float64
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) area() float64 {
	return t.base * t.height / 2
}

func printShapeArea(s shape) {
	fmt.Println("The area of given shape", s.area())
}

func main() {
	s := square{sideLength: 3}
	t := triangle{base: 10, height: 3}

	printShapeArea(s)
	printShapeArea(t)
}
