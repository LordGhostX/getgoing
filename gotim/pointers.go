package main

import "fmt"

func changeValue(str *string) {
	*str = "changed"
}

func changeValue2(str string) {
	str = "changed"
}

func Pointers() {
	x := 7
	fmt.Println(x, &x)

	y := &x
	*y = 8
	fmt.Println(x, *y, y)

	str, str2 := "hello", "hello2"
	fmt.Println(str, str2)
	changeValue(&str)
	changeValue2(str2)
	fmt.Println(str, str2)
}
