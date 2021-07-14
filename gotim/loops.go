package main

import "fmt"

func Loops() {
	x := 3
	for x < 10 {
		fmt.Println(x)
		x++
	}

	for i := 3; i < 10; i++ {
		fmt.Println(i)
	}

	//for {
	//
	//}

	for i := 0; i < 1000; i++ {
		if i != 0 && i%11 == 0 && i%13 == 0 {
			fmt.Println(i)
			continue
		}

		if i == 800 {
			break
		}
	}
}
