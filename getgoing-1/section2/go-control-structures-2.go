package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	for i := 0; i < 100; i++ {
		if i == 50 {
			fmt.Println("Finally Out!")
			break
		}
	}

	day := "Fri"
	switch day {
	case "Mon":
		fmt.Println("Today is Monday")
  case "Tue", "Thur":
    fmt.Println("Today begins with T")
  case "Fri", "Sat", "Sun":
    fmt.Println("Yayyy!!! It's Weekend")
  default:
    fmt.Println("Some Other Day Instead")
	}

	switch {
	case day == "Mon":
		fmt.Println("Today is Monday")
  case day == "Tue" || day == "Thur":
    fmt.Println("Today begins with T")
  case day == "Fri" || day == "Sat" || day == "Sun":
    fmt.Println("Yayyy!!! It's Weekend")
  default:
    fmt.Println("Some Other Day Instead")
	}
}
