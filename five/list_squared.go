package five

import "math"

// ListSquared find all integers between m and n whose sum of squared divisors is itself a square
func ListSquared(m, n int) [][]int {
	results := [][]int{}
	for i := m; i < n; i++ {
		squaredSum := GetSquaredSumDivisors(i)
		root := int(math.Sqrt(float64(squaredSum)))
		if root*root == squaredSum {
			results = append(results, []int{i, squaredSum})
		}
	}
	return results
}

// GetSquaredSumDivisors returns the divisors of an integer
func GetSquaredSumDivisors(integer int) int {
	squaredSum := 0
	for i := 1; i <= integer; i++ {
		div := integer / i
		if integer == div*i {
			squaredSum += int(math.Pow(float64(i), 2))
		}
	}
	return squaredSum
}
