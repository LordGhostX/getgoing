package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Controls() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum *= 2
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum *= 2
	}
	fmt.Println(sum)

	// loop forever
	//for {
	//	fmt.Println("Print Forever!")
	//}

	if true {
		fmt.Println("Hello World!")
	}

	rand.Seed(time.Now().UnixNano())
	if v := rand.Int() % 100; v < 50 {
		fmt.Println("Random Number is Small")
	} else {
		fmt.Println("Random Number is big")
	}

	switch 3 {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println("Something else!")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}

	defer func() {
		fmt.Println("Ending the code here!")
	}()
	defer fmt.Println("Hello World Again!")
}
