package main

import (
	"fmt"
	"sync"
	"time"
)

func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Something!")
	}
}

func main() {
	//go heavy()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		fmt.Println("Hello")
		wg.Done()
	}()
	go func() {
		fmt.Println("World")
		wg.Done()
	}()

	fmt.Println("Fin")
	wg.Wait()
	fmt.Println("Exit!")
	//select {}
}
