package main

import "fmt"

func Arrays() {
	var arr [5]int
	arr[0] = 100
	arr[4] = 200
	fmt.Println(arr)
	fmt.Println(arr[0])

	arr2 := [3]int{1, 2 , 3}
	fmt.Println(arr2)
	fmt.Println(len(arr2))

	sum := 0
	for i := 0; i < len(arr2); i++ {
		sum += arr2[i]
	}
	fmt.Println(sum)

	arr3 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(arr3)
	fmt.Println(arr3[0])
	fmt.Println(arr3[0][2])
}
