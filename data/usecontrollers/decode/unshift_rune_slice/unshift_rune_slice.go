package unshift_rune_slice

const LIMITER = 1_112_064

// UnshiftRuneSlice shifts every rune os the slice of runes
func UnshiftRuneSlice(rx []rune, shifter int) []rune {
	var nrx []rune

	for _, r := range rx {
		rint := int(r)
		sftrOp := rint - shifter

		if sftrOp >= 0 {
			nrc := rune(sftrOp)
			nrx = append(nrx, nrc)
		} else {
			newSftrOp := LIMITER + sftrOp
			nrc := rune(newSftrOp)
			nrx = append(nrx, nrc)
		}
	}

	return nrx
}
