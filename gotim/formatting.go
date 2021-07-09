package main

import "fmt"

func Format() {
	fmt.Printf("%v %t %T\n", 123, true, 123.456)
	fmt.Printf("%b %o %d %x %X\n", 34, 67, 54, 943, 943)
	fmt.Printf("%e %f %F %g\n", 3.67856576897867, 5.21387678, 9.56838693, 5.9395789)
	fmt.Printf("%s %q\n", "1234", "abcd")
	fmt.Printf("%20f %.2f %9.2f %9.f\n", 3.568, 5.789033, 4.5943, 2.689)
	fmt.Printf("%10v %10X\n", 678, 667)
	fmt.Printf("%-20s is cool\n", "Golang")
	fmt.Printf("%07d %0d\n", 123, 456)

	greeting := fmt.Sprintf("%q World", "Hello")
	fmt.Println(greeting)
}
