package eight

// ReverseString the string provided
func ReverseString(word string) string {
	var str string
	for i := len(word); i > 0; i-- {
		str += string(word[i-1])
	}
	return str
}
