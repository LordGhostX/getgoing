package main

import "fmt"

func main() {
	arr2 := []int{1, 2, 3, 4}
	arr3 := []string{"hello", "world", "everyone"}

	arr2 = append(arr2, 5, 6, 7)
	fmt.Println(arr2, arr3)
}
