package main

import (
	"fmt"
	"os"
)

func main() {
	c := make(chan int)
	q := make(chan int)

	go func() {
		c <- 1
	}()

	for {
		select {
		case <-c:
			fmt.Println("Received Message!")
			go func() {
				q <- 1
			}()
		case <-q:
			fmt.Println("Exiting App!")
			os.Exit(0)
		}
	}
}
