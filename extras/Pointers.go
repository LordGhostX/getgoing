package main

import "fmt"

func swap(var1, var2 *int) {
	var temp int
	temp = *var2
	*var2 = *var1
	*var1 = temp
}

func main() {
	var1 := 5
	var2 := 10

	fmt.Println(var1, var2)
	swap(&var1, &var2)
	fmt.Println(var1, var2)
}
