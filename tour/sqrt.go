package main

import (
	"math"
)

func Sqrt(x float64) float64 {
	z := x / 2
	p := z
	for {
		z -= (z*z - x) / (2 * z)
		if math.Abs(x-p*p) <= math.Abs(x-z*z) {
			return p
		} else {
			p = z
		}
	}
}
