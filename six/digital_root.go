package six

// DigitalRoot calculates the sum of the integer of a number recursively
func DigitalRoot(n int) int {
	if n < 10 {
		return n
	}
	newN := n / 10
	val := n % 10
	return DigitalRoot(newN + val)
}
