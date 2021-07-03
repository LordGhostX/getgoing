package main
import (
  "fmt"
  "time"
)

func heavy() {
  for {
    time.Sleep(time.Second * 1)
    fmt.Println("Heavy")
  }
}

func superHeavy() {
  for {
    time.Sleep(time.Second * 3)
    fmt.Println("Super Heavy")
  }
}

func main() {
  fmt.Println("Start...")
  go heavy()
  go superHeavy()
  select {}
}
