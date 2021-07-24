package main

import "fmt"

func Map() {
	var mp map[string]int
	fmt.Println(mp)

	var mp2 = map[string]int{
		"apple":  10,
		"pear":   5,
		"orange": 7,
	}
	fmt.Println(mp2)
	fmt.Println(mp2["apple"])
	fmt.Println(len(mp2))

	mp2["apple"] = 100
	mp2["banana"] = 200
	delete(mp2, "pear")
	fmt.Println(mp2)

	mp3 := make(map[string]int)
	fmt.Println(mp3)

	mp3["orange"] = 10
	val, ok := mp3["apple"]
	fmt.Println(val, ok)

	val, ok = mp3["orange"]
	fmt.Println(val, ok)
}
