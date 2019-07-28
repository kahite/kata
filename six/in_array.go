package six

import (
	"sort"
	"strings"
)

// InArray returns the sorted list of strings in array1 that exist in array2
func InArray(array1 []string, array2 []string) []string {
	possibleResults := sort.StringSlice(array1)
	possibleResults.Sort()

	results := []string{}

	for _, a := range possibleResults {
		if len(results) != 0 && len(InArray([]string{a}, results)) != 0 {
			continue
		}

		for _, b := range array2 {
			if strings.Contains(b, a) {
				results = append(results, a)
				break
			}
		}
	}

	return results
}
