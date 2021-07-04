package main

import (
	"fmt"
	"rsc.io/quote"
)

func Quote() {
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Opt())
}
