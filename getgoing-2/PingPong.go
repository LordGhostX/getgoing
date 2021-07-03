package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ping := make(chan int)
	pong := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(10)

	go pingBall(ping, pong, wg)
	go pongBall(ping, pong, wg)

	pong <- 1

	wg.Wait()
	fmt.Println("End!")
}

func pingBall(ping chan int, pong chan int, wg *sync.WaitGroup) {
	for {
		<- pong
		wg.Done()
		fmt.Println("Ping Ball")
		time.Sleep(time.Second)
		ping <- 1
	}
}

func pongBall(ping chan int, pong chan int, wg *sync.WaitGroup) {
	for {
		<- ping
		wg.Done()
		fmt.Println("Pong Ball")
		time.Sleep(time.Second)
		pong <- 1
	}
}
