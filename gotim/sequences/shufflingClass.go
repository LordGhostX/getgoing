package sequences

func shuffleClass(class []int, n int) []int {
	if n < 0 {
		n *= -1
	}
	return append(class[n:], class[:n]...)
}
