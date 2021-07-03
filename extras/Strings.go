package main

import (
	"fmt"
	"strings"
)

func main() {
	var var1 string
	var1 = "hello world!"
	fmt.Println(var1)

	var1 = "another thing!"
	fmt.Println(var1)

	fmt.Println(strings.Contains(var1, "another"))

	fmt.Println(strings.ReplaceAll(var1, "t", "T"))

	fmt.Println(strings.Split(var1, " "))
}
