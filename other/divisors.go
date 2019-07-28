package other

// GetDivisors returns the divisors of an integer
func GetDivisors(integer int) []int {
	divisors := []int{}
	for i := 1; i <= integer; i++ {
		div := integer / i
		if integer == div*i {
			divisors = append(divisors, i)
		}
	}
	return divisors
}
