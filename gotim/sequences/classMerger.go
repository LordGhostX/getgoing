package sequences

func mergeClasses(a, b []int) []int {
	for idB := range b {
		if b[idB] > a[len(a) - 1] {
			a = append(a, b[idB])
			continue
		}
		for idA := range a {
			if b[idB] <= a[idA] {
				a = append(a[:idA], append([]int{b[idB]}, a[idA:]...)...)
				break
			}
		}
	}

	return a
}
