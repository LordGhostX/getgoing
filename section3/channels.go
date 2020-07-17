package main
import (
  "fmt"
)

func main() {
  var c chan int
  c2 := make(chan int)

  fmt.Println(c)
  fmt.Println(c2)

  go func() {
    c2 <- 10
  }()

  val := <- c2
  fmt.Println(val)

  go func() {
    c2 <- 20
  }()

  val = <- c2
  fmt.Println(val)
}
