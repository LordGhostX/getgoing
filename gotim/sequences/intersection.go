package sequences

import (
	"strings"
)

func findIntersection(a, b string) []string {
	var intersection []string
	aMap := make(map[string]interface{})
	for _, i := range strings.Split(a, ", ") {
		aMap[i] = nil
	}
	for _, i := range strings.Split(b, ", ") {
		if _, j := aMap[i]; j {
			intersection = append(intersection, i)
		}
	}

	if len(intersection) == 0 {
		return []string{"false"}
	} else {
		return intersection
	}
}
