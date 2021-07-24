package main

import "fmt"

type student struct {
	name   string
	grades []int
	age    int
}

func (s student) getAge() int {
	return s.age
}

func (s *student) setAge(age int) {
	s.age = age
}

func (s student) getAvgGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func (s student) getMaxGrade() int {
	currentMax := 0
	for _, v := range s.grades {
		if v > currentMax {
			currentMax = v
		}
	}
	return currentMax
}

func Methods() {
	s1 := student{"LordGhostX", []int{70, 62, 65, 40}, 19}
	fmt.Println(s1.getAge())

	s1.setAge(20)
	fmt.Println(s1.getAge())

	fmt.Println(s1.getAvgGrade())

	s2 := student{"Solomon", []int{80, 70, 75, 85, 68}, 21}
	fmt.Println(s2.getAvgGrade())

	fmt.Println(s1.getMaxGrade())
	fmt.Println(s2.getMaxGrade())
}
