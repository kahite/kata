package six

import (
	"codewarrior/kata/eight"
	"math"
	"strconv"
)

// Revrot reverses or rotate string chunks
func Revrot(s string, n int) string {
	if len(s) == 0 || n <= 0 {
		return ""
	}

	result := ""
	chunk := ""
	ind := 0
	for index, c := range s {
		ind = index
		if ind%n == 0 && ind != 0 {
			result += ProcessChunk(chunk)
			chunk = ""
		}

		chunk += string(c)
	}
	if ind%n == n-1 && ind != 0 {
		result += ProcessChunk(chunk)
	}

	return result
}

// ProcessChunk ...
func ProcessChunk(chunk string) string {
	intVal, _ := strconv.Atoi(chunk)
	if GetCubicDigits(intVal)%2 == 0 {
		return eight.ReverseString(chunk)
	}
	return RotateChunkToLeft(chunk)
}

// GetCubicDigits ...
func GetCubicDigits(n int) int {
	if n == 0 {
		return 0
	}
	return int(math.Pow(float64(n%10), 3)) + GetCubicDigits(n/10)
}

// RotateChunkToLeft brings the last character in the end
func RotateChunkToLeft(chunk string) string {
	var str string
	for i := 1; i < len(chunk); i++ {
		str += string(chunk[i])
	}
	str += string(chunk[0])
	return str
}
