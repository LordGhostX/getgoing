package sequences

import "testing"

func TestIndexFinding(t *testing.T) {
	i := findIndex([]int{0, 8, -2, 5, 0}, 0)
	if i != [2]int{1, 2} {
		t.Fail()
	}
}

func TestIndexFindingReturnsNothingWhenValNotFound(t *testing.T) {
	i := findIndex([]int{0, 8, -2, 5, 0}, 1)
	if i != [2]int{-1, -1} {
		t.Fail()
	}
}
