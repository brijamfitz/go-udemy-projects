package main

import (
	"fmt"
	"math"
)

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
	t := triangle{
		height: 3,
		base:   5,
	}

	s := square{
		sideLength: 7,
	}

	at := printArea(t)
	as := printArea(s)

	fmt.Println("Area of a triangle:", at, "feet")
	fmt.Println("Aread of a square:", as, "feet")
}

func printArea(s shape) float64 {
	return s.getArea()
}

func (t triangle) getArea() float64 {
	return (t.base * t.height) / 2
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}
