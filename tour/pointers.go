package main

import "fmt"

type Vector struct {
	X int
	Y int
	Z int
}

type Vertex struct {
	X, Y int
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	index, fib, fib2 := 0, 0, 1
	return func() int {
		index++
		if index <= 2 {
			return index - 1
		}
		fib, fib2 = fib2, fib2 + fib
		return fib2
	}
}

func main() {
	i := 42
	p := &i

	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(i)

	*p = 30
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(i)

	v := Vector{1, 2, 3}
	fmt.Println(v)

	v2 := &v
	v2.X = 10
	fmt.Println(v2.X, v2.Y, v2.Z)

	v3 := Vertex{
		X: 10,
		Y: 15,
	}
	fmt.Println(v3)

	v4 := &Vertex{3, 5}
	fmt.Println(*v4)

	var arr [10]int
	arr[0] = 5
	arr[3] = 3
	fmt.Println(arr)

	arr2 := []string{"Hello", "World", "!"}
	fmt.Println(arr2)

	fmt.Println(arr[0:4])
	fmt.Println(arr2[1:])

	slice := arr[:4]
	slice[1] = 2
	fmt.Println(slice)
	fmt.Println(arr)

	fmt.Println(len(arr))
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	arr3 := make([]int, 5)
	fmt.Println(arr3)

	arr3 = append(arr3, 1, 2, 3)
	fmt.Println(arr3)

	for i, v := range arr3[5:] {
		fmt.Printf("%v * %v = %v\n", i, v, i*v)
	}

	m := make(map[string]Vertex)
	m["p1"] = Vertex{1, 3}
	fmt.Println(m["p1"])

	m2 := map[string]Vertex {
		"p1": {1, 3},
		"p2": {10, 3},
	}
	fmt.Println(m2)
	fmt.Println(m2["p1"])

	m2["p1"] = Vertex{3, 4}
	fmt.Println(m2["p1"])

	delete(m2, "p1")
	fmt.Println(m2)

	elem, ok := m2["p2"]
	fmt.Println(elem, ok)

	elem, ok = m2["p3"]
	fmt.Println(elem, ok)

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}