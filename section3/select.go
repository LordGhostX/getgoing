package main

import (
	"fmt"
	"sync"
)

func test(c chan int, quits chan struct{}) {
	select {
	case <- c:
		fmt.Println("Received")
	case <- quits:
		fmt.Println("Quit")
		// os.Exit(0)
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	c := make(chan int, 2)
	quits := make(chan struct{})

	go func() {
		c <- 1
		quits <- struct{}{}
	}()

	go test(c, quits)
	wg.Wait()
	fmt.Println("Hello")
}
