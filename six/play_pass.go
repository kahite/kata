package six

import (
	"sort"
	"strconv"
	"strings"
)

// PlayPass ...
// 1. shift each letter by a given number but the transformed letter must be a letter (circular shift),
// 2. replace each digit by its complement to 9,
// 3. keep such as non alphabetic and non digit characters,
// 4. downcase each letter in odd position, upcase each letter in even position (the first character is in position 0),
// 5. reverse the whole result.
func PlayPass(s string, n int) string {
	alphabet := sort.StringSlice([]string{
		"a", "b", "c", "d", "e",
		"f", "g", "h", "i", "j",
		"k", "l", "m", "n", "o",
		"p", "q", "r", "s", "t",
		"u", "v", "w", "x", "y", "z",
	})
	result := ""

	s = strings.ToLower(s)
	for index, c := range s {
		newValue := string(c)
		if i, err := strconv.Atoi(newValue); err == nil {
			newValue = strconv.Itoa(9 - i)
		} else if alphabetIndex := GetIndex(alphabet, newValue); alphabetIndex >= 0 {
			newValue = alphabet[(alphabetIndex+n)%26]
			if index%2 == 0 {
				newValue = strings.ToUpper(newValue)
			}
		}

		result = newValue + result
	}

	return result
}

// GetIndex returns the index of a string in slice or -1 if it does not exist
func GetIndex(a []string, x string) int {
	for index, n := range a {
		if x == n {
			return index
		}
	}
	return -1
}
