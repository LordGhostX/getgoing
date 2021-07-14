package main

import "fmt"

func Switches() {
	val := -1

	switch val {
	case 1, -1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("not a case")
	}

	switch {
	case val > 0:
		fmt.Println("greater than 0")
	case val < 0:
		fmt.Println("less than 0")
	default:
		fmt.Println("zero")
	}
}
