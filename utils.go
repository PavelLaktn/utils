package utils

// Contains a slice a specific string or not
func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

// Contains a slice a specific int or not
func ContainsInt(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
