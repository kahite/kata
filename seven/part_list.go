package seven

// PartList _
func PartList(arr []string) string {
	ret := ""
	length := len(arr)

	for h := 0; h < length-1; h++ {
		ret += "("
		for i := 0; i < length; i++ {
			if i != 0 {
				ret += " "
			}
			ret += arr[i]
			if i == h {
				ret += ","
				for j := i + 1; j < length; j++ {
					i++
					ret += " " + arr[j]
				}
			}
		}
		ret += ")"
	}

	return ret
}
