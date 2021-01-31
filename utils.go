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
