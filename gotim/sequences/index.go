package sequences

func findIndex(nums []int, val int) [2]int {
	valCount, valPos := 0, 0
	for _, j := range nums {
		if val > j {
			valPos++
		}
		if j == val {
			valCount++
		}
	}
	if valCount == 0 {
		return [2]int{-1, -1}
	} else {
		return [2]int{valPos, valPos + valCount - 1}
	}
}
