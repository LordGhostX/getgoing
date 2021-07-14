package main

import (
	"fmt"
)

func Slices() {
	x := [5]int{1, 2, 3, 4, 5}
	s := x[1:3]

	fmt.Printf("%T %T\n", x, s)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(s[:cap(s)])

	s = append(s, 6, 7, 8, 9, 10)
	fmt.Println(s)

	arr := make([]int, 5)
	fmt.Println(arr)
	fmt.Printf("%T\n", arr)
}
