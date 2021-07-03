package main

import (
	"fmt"
)

type Vehicle struct {
	Type string
}

func main() {
	var c chan int
	c = make(chan int)
	fmt.Println(c)

	c2 := make(chan int)
	fmt.Println(c2)

	go func() {
		c2 <- 1
	}()
	fmt.Println(<-c2)

	go func() {
		c2 <- 2
	}()
	fmt.Println(<-c2)

	c3 := make(chan int, 3)
	go func() {
		c3 <- 3
		c3 <- 4
		c3 <- 5
		c3 <- 6
		close(c3)
	}()

	for i := range c3 {
		fmt.Println(i)
	}

	c4 := make(chan *Vehicle, 3)
	go func() {
		c4 <- &Vehicle{"Plane"}
		c4 <- &Vehicle{"Train"}
		c4 <- &Vehicle{"Bicycle"}
		c4 <- &Vehicle{"Skateboard"}
		close(c4)
	}()

	for i := range c4 {
		fmt.Println(i.Type)
	}

	c5 := make(chan int)
	fmt.Println(<-c5)
}