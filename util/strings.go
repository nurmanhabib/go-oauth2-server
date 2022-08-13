package util

// SliceContains is a function to make sure a string is in the slice of string.
func SliceContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
