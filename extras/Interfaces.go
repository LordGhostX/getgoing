package main

import "fmt"

type Car interface {
	Drive()
	Stop()
	Print()
}

type Lambo struct {
	Model string
}

func NewLambo(model string) Car {
	return &Lambo{model}
}

func (l Lambo) Print() {
	fmt.Println("Printing", l.Model)
}

func (l Lambo) Drive() {
	fmt.Println("Driving", l.Model)
}

func (l Lambo) Stop() {
	fmt.Println("Stopping", l.Model)
}

func PrintAnything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	lambo := NewLambo("Gallardo")
	lambo.Print()
	lambo.Drive()
	lambo.Stop()

	PrintAnything(3.142)
	PrintAnything(10)
	PrintAnything("Something!")
}
