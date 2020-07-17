package main

import (
	"fmt"
)

type Car interface {
	Drive()
	Stop()
}

type Lambo struct {
	LamboModel string
}

type Ford struct {
	FordModel string
}

func (l *Lambo) Drive() {
  fmt.Println("Lambo on the move")
  fmt.Println(l.LamboModel)
}

func (f *Ford) Drive() {
  fmt.Println("Ford on the move")
  fmt.Println(f.FordModel)
}

func Anything(anything interface{}) {
  fmt.Println(anything)
}

func main() {
	l := Lambo{"Gallardo"}
  f := Ford{"Ultra"}

  l.Drive()
  f.Drive()

  Anything("Hello World")
  Anything(12.6748)
  Anything(15)
}
