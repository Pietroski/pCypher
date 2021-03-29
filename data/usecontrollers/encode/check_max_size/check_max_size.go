package check_max_size

// CheckMaxSize checks the maximum size of a rune slice
func CheckMaxSize(iasSize int, asSize int, injectionFactor int) bool {
	ms := MaxSize(iasSize, injectionFactor)
	if asSize != ms {
		return false
	}

	return true
}

// MaxSize checks the maximum size of a rune slice
func MaxSize(rsSize int, ijf int) int {
	maxSize := (rsSize + 1) * ijf + rsSize

	return maxSize
}