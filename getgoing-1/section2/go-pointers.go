package main

import (
	"fmt"
)

func swap(var1, var2 *int) {
	var temp int
	temp = *var2
	*var2 = *var1
	*var1 = temp
}

func main() {
	var1 := "my variable"
	pointer1 := &var1

	fmt.Println(var1)
	fmt.Println(pointer1)
	fmt.Println(*pointer1)

	var (
		m1 = 10
		m2 = 5
	)

	fmt.Println(m1, m2)
	swap(&m1, &m2)
	fmt.Println(m1, m2)
}
