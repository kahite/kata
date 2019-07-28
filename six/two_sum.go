package six

// TwoSum retrieves the 2 numbers that matches the sum
func TwoSum(numbers []int, target int) [2]int {
	numMap := make(map[int]int)
	for ind, n := range numbers {
		numMap[n] = ind
	}

	for ind2, n := range numbers {
		diff := target - n
		if _, ok := numMap[diff]; ok {
			return [2]int{ind2, numMap[diff]}
		}
	}

	return [2]int{}
}
