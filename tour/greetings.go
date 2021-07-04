package main

import (
	"fmt"
	"golang/greetings"
	"log"
)

func Greet() {
	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	message, err := greetings.Hello("LordGhostX")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(message)

	names := []string{
		"LordGhostX",
		"Solomon",
		"John",
		"Sharon",
		"Samuel",
	}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, message := range messages {
		fmt.Println(message)
	}
}
