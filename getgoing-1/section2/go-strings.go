package main
import (
  "fmt"
  "strings"
)

func main() {
  var var1 string
  var1 = "my name"

  var2 := "my second name"

  fmt.Println(var1)
  fmt.Println(var2)
  fmt.Println(strings.Contains(var2, "name"))
  fmt.Println(strings.ReplaceAll(var1, "m", "n"))
  fmt.Println(strings.Split(var2, "m"))
  fmt.Println(var1 + " " + var2)
  fmt.Println(var1, var2)
}
