package main
import (
  "fmt"
)

type Car struct {
  Model string
}

func main() {
  c := make(chan *Car, 10)

  go func() {
    c <- &Car{"1"}
    c <- &Car{"2"}
    c <- &Car{"3"}
    c <- &Car{"4"}
    c <- &Car{"5"}
    close(c)
  }()

  for i := range c {
    fmt.Println(i.Model)
  }
}
