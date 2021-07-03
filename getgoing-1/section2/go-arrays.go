package main
import (
  "fmt"
)

func todo() {
  arr1 := []int {1, 2, 3, 4}
  arr2 := []string {"my", "name", "is"}

  arr2 = append(arr2, "ghost", "!!!")

  fmt.Println(arr1)
  fmt.Println(arr2)
}

func main() {
  todo()
}
