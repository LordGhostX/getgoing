package main
import (
  "fmt"
)

func main() {
  fmt.Println("Hello World!")
  age := 18

  if age > 20 {
    fmt.Println("You are an adult")
  } else if age == 18 {
    fmt.Println("Welcome to adulthood")
  } else {
    fmt.Println("You are a minor")
  }

  if (age == 18 || age > 15) && age < 20 {
    fmt.Println("You are greater than 15")
  }

  for i := 0; i < 10; i++ {
    fmt.Println(i)
  }

  counter := 0
  for {
    counter += 1
    fmt.Println("Printing Infinitely...", counter)
    if counter == 20 {
      break
    }
  }

  arr := []string {"Hello", "World", "!", "Once", "Again"}
  for k, v := range arr {
    fmt.Println(k, v)
  }

  mymap := make(map[string]interface{})
  mymap["name"] = "Ghost"
  mymap["age"] = 18

  for k, v := range mymap {
    fmt.Printf("Key: %s and Value %s", k, v)
  }
}
