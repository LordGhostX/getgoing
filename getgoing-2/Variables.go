package main

import "fmt"

func main() {
	var var1 int
	var1 = 10
	fmt.Println(var1)

	var (
		var2 = 5
		var3 = 8
	)
	fmt.Println(var2 + var3)

	var var4 int32
	var var5 int64
	fmt.Println(int64(var4) + var5)

	var6 := 4
	var7 := 10
	fmt.Println(var6 + var7)

	var8, var9 := 10, 6
	fmt.Println(var8 + var9)
}
