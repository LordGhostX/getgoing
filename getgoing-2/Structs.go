package main

import "fmt"

type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) Drive() {
	fmt.Println(c.Name, "is driving!")
}

func (c Car) GetName() string {
	return c.Name
}

func main() {
	struct1 := Car{}
	fmt.Println(struct1)

	struct2 := Car{"Lambo", 3, 12089}
	fmt.Println(struct2)

	struct3 := Car{
		Name:    "Honda",
		Age:     9,
		ModelNo: 6748,
	}
	fmt.Println(struct3)

	struct2.Print()
	struct3.Drive()
	fmt.Println(struct3.GetName())
}
