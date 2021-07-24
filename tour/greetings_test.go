package main

import (
	"golang/greetings"
	"strings"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "LordGhostX"
	message, err := greetings.Hello(name)
	if err != nil || !strings.Contains(message, name) {
		t.Fatalf("Greeting should contain %v", name)
	}
}

func TestHelloEmpty(t *testing.T) {
	message, err := greetings.Hello("")
	if err == nil || message != "" {
		t.Fail()
	}
}

func TestHellos(t *testing.T) {
	names := []string{
		"LordGhostX",
		"Solomon",
		"John",
	}
	messages, err := greetings.Hellos(names)
	if err != nil || len(messages) != len(names) {
		t.Fatalf("Greetings should be %v in count", len(names))
	}
}

func TestHellosEmpty(t *testing.T) {
	names := []string{""}
	messages, err := greetings.Hellos(names)
	if err == nil || messages != nil {
		t.Error("Greetings should not be empty")
	}
}
