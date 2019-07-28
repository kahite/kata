package six

// SplitStrings splits the param in 2 characters strings
func SplitStrings(str string) []string {
	return recSplit(str, []string{})
}

func recSplit(str string, slice []string) []string {
	length := len(str)
	if length == 0 {
		return slice
	} else if length == 1 {
		return append(slice, str+"_")
	} else {
		return recSplit(str[2:], append(slice, str[:2]))
	}
}
