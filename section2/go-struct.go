package main

import (
  "fmt"
)

type Car struct {
  Name  string
  Age int
  ModelNo int
}

func (c Car) Print() {
  fmt.Println(c)
}

func (c Car) Drive() {
  fmt.Println("Driving...")
}

func (c Car) GetName() string {
  return c.Name
}

func main() {
  car1 := Car {"Ford", 10, 10001}
  car2 := Car {
    Name: "Toyota",
    Age: 15,
    ModelNo: 23450,
  }

  fmt.Println(car1)
  fmt.Println(car2)

  car1.Print()
  car2.Drive()

  fmt.Println(car1.GetName())
}
