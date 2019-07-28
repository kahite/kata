package seven

import "codewarrior/kata/six"

// ContainAllRots check if a slice contains all the rotations of a string
func ContainAllRots(strng string, arr []string) bool {
	if strng == "" {
		return true
	}

	rotatedStrings := []string{}

	rotatedStrings = append(rotatedStrings, strng)
	for i := 1; i < len(strng); i++ {
		rotatedStrings = append(rotatedStrings, six.RotateChunkToLeft(rotatedStrings[i-1]))
	}

	for _, s := range rotatedStrings {
		if six.GetIndex(arr, s) == -1 {
			return false
		}
	}

	return true
}
