package main

import "fmt"

func Booleans() {
	var1 := true
	var2 := false

	fmt.Printf("%t %t\n", var1, var2)

	if 5 < 10 && 10 > 5 {
		fmt.Println("Chained!")
	}

	if 5 > 10 || 10 > 5 {
		fmt.Println("Chained!!")
	}
}
