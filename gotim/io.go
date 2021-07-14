package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func IO() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter something: ")
	scanner.Scan()
	fmt.Printf("You typed %q\n", scanner.Text())

	var text string
	fmt.Print("Enter something: ")
	fmt.Scanf("%s", &text)
	fmt.Printf("You typed %q\n", text)

	fmt.Println(strconv.ParseInt("100",10, 64))
	fmt.Println(strconv.Atoi("100"))
}