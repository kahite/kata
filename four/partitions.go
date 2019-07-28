package four

// Partitions returns the number of integer partitions of n
func Partitions(n int) int {
	if n < 4 {
		return n
	}

	result := 0

	result += Partitions(n - 1)                 // Numbers with 1
	result += Partitions(n-2) - Partitions(n-3) // Numbers with 2
	if n >= 6 {
		result += Partitions(n-3) - Partitions(n-4) //Numbers with 3
	}
	if n >= 7 {
		result += Partitions(n-6) - Partitions(n-5) //Numbers with 3 correctio 2 sur 3
	}
	if n >= 8 {
		result += Partitions(n-4) - Partitions(n-5) // Numbers with 4
		result += Partitions(n-7) - Partitions(n-6) // Correction 2 sur 4
	}

	return result + 1
}
