package main

import "fmt"

type point struct {
	x int
	y int
}

type circle struct {
	radius float64
	center *point
}

func changeX(pt *point, x int) {
	pt.x = x
}

func changeY(pt *point, y int) {
	pt.y = y
}

func Structs() {
	p1 := point{10, 15}

	fmt.Println(p1)
	p1.x = 20
	fmt.Println(p1)

	p2 := point{x: 5}
	fmt.Println(p2)

	changeX(&p2, 10)
	fmt.Println(p2)
	changeX(&p2, 15)
	fmt.Println(p2)

	p3 := &point{y: 5}
	fmt.Println(p3)
	changeY(p3, 10)
	fmt.Println(*p3)

	c1 := circle{4.65, &point{5, 8}}
	fmt.Println(c1)
	fmt.Println(c1.center.y)
}