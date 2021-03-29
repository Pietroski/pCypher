package check_min_size

// CheckMinSize checks the maximum size of a rune slice
func CheckMinSize(initialSliceSize int, finalSliceSize int, injectionFactor int) bool {
	ms := MinSize(initialSliceSize, injectionFactor)
	if finalSliceSize != ms {
		return false
	}

	return true
}

// MinSize checks the minimum size of a rune slice
func MinSize(rsSize int, injf int) int {
	letters := 0
	litters := 0

	var checker int

	minSize := rsSize - (2 * injf)

	for i := 0; i < minSize && checker < minSize; i++ {
		letters++
		litters += injf

		checker = letters + litters
	}

	return letters
}
