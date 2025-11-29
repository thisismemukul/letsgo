package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Printf("%T area: %.2f\n", s, s.area())
}

// runAreaDemo demonstrates calculating areas for different shapes.
func runAreaDemo() {
	shapes := []shape{
		circle{radius: 10},
		rectangle{width: 10, height: 10},
		triangle{base: 10, height: 5},
	}

	for _, s := range shapes {
		printArea(s)
	}
}
