package remove_random_runes

// RemoveRandomRunes from a slice of runes
func RemoveRandomRunes(rs []rune, injectionFactor int) []rune {
	rsSize := len(rs)
	var nrs = make([]rune, 0, rsSize)

	index := injectionFactor
	for index < rsSize {
		v := rs[index]
		nrs = append(nrs, v)

		index += injectionFactor + 1
	}

	return nrs
}
