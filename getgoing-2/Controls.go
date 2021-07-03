package main

import "fmt"

func main() {
	age := 20

	if age >= 18 {
		fmt.Println("Old enough to be jailed!")
	} else if age >= 5 && age < 18 {
		fmt.Println("A minor!")
	} else {
		fmt.Println("A toddler!")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	arr := []string{"my", "name", "is", "Lord", "Ghost", "X!"}
	for i, j := range arr {
		fmt.Println(i, j, arr[i])
	}

	myMap := make(map[string]interface{})
	myMap["name"] = "LordGhostX!"
	myMap["age"] = -999

	for k, v := range myMap {
		fmt.Println(k, v)
	}

	i := 0
	for {
		if i >= 10 {
			break
		}

		fmt.Println("looping infinitely!")
		i++
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	day := "Friday"
	switch day {
	case "Friday":
		fmt.Println("TGIF!")
	case "Monday", "Tuesday", "wednesday", "Thursday":
		fmt.Println("Boring!")
	default:
		fmt.Println("Default")
	}
}
