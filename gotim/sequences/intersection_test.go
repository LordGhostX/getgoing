package sequences

import "testing"

func TestFindIntersection(t *testing.T) {
	intersection := findIntersection("1, 3, 4, 7, 13", "1, 2, 4, 13, 15")
	for i := range intersection {
		if intersection[i] != []string{"1", "4", "13"}[i] {
			t.Fail()
		}
	}

	intersection = findIntersection("1, 3, 9, 10, 17, 18", "1, 4, 9, 10")
	for i := range intersection {
		if intersection[i] != []string{"1", "9", "10"}[i] {
			t.Fail()
		}
	}
}

func TestFindIntersectionReturnsFalseIfInvalid(t *testing.T) {
	intersection := findIntersection("1, 2, 3", "4, 5, 6")
	if intersection[0] != []string{"false"}[0] || len(intersection) != 1 {
		t.Fail()
	}
}
