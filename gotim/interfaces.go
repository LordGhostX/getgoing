package main

import "fmt"

type shape interface {
	area() float64
}

type triangle struct {
	base float64
	height float64
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

type square struct {
	size float64
}

func (s square) area() float64 {
	return s.size * s.size
}

func getArea(s shape) float64 {
	return s.area()
}

func Interface() {
	shapes := []shape{triangle{10, 5}, square{10}}

	for _, s := range shapes {
		fmt.Println(s.area())
		fmt.Println(getArea(s))
	}
}