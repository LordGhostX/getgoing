package main

import "fmt"

func test() {
	fmt.Println("Testing!!!")
}

func printInt(x int) {
	fmt.Println(x)
}

func add(x, y int) int {
	return  x + y
}

func addSub(x, y int) (int, int) {
	return x + y, x - y
}

func addSub2(x, y int) (a, b int) {
	defer fmt.Println("hello!!!")

	a, b = x + y, x - y
	return
}

func test2(myFunc func(int) int, x int) {
	fmt.Println(myFunc(x))
}

func Functions() {
	test()
	printInt(10)
	printInt(5)
	fmt.Println(add(5, 8))
	fmt.Println(addSub(10, 3))
	fmt.Println(addSub2(10, 3))

	x := test
	x()
	y := printInt
	y(10)

	func() {
		fmt.Println("Hello!")
	}()

	t2 := func(x int) int {
		return x * -1
	}
	fmt.Println(t2(8))
	test2(t2, 10)
}
