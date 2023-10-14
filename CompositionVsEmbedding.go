package main

import (
	"fmt"
)

type car struct {
	name  string
	price int
}

type MyNil interface {
}

func (c *car) feature() {
	fmt.Printf("This is a %v with price %v\n", c.name, c.price)
}

func (c *car) sayName() {
	fmt.Printf("This is a %v\n", c.name)
}

func (c *car) cost() {
	fmt.Printf("The price is %d\n", c.price)
}

type bmw struct {
	*car
	average float32
}

func (b *bmw) cost() {
	fmt.Printf("The price of %v is %d\n", b.name, 10000000)
}

type bentley struct {
	average float32
}
