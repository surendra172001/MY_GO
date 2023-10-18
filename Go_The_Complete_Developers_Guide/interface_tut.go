package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (s *Square) area() float64 {
	return s.side * s.side
}

func (s *Square) circum() float64 {
	return 4 * s.side
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) circum() float64 {
	return 2 * math.Pi * c.radius
}

type shape interface {
	area() float64
	circum() float64
}

func printGeoInfo(s *shape) {
	fmt.Printf("Area: %.2f, Circumference: %.2f\n", (*s).area(), (*s).circum())
}

func Interface_tut() {
	var shapes = []shape{
		&Circle{radius: 2},
		&Square{side: 2},
		&Circle{radius: 2.5},
	}

	for _, v := range shapes {
		printGeoInfo(&v)
		fmt.Println("...")
	}
}
