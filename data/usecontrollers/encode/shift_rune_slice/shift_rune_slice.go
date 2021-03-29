package shift_rune_slice

const LIMITER = 1_112_064

// ShiftRuneSlice shifts every rune os the slice of runes
func ShiftRuneSlice(rx []rune, shifter int) []rune {
	var nrx []rune

	for _, r := range rx {
		rint := int(r)
		sftrOp := rint + shifter

		if sftrOp < LIMITER {
			nrc := rune(sftrOp)
			nrx = append(nrx, nrc)
		} else {
			newSftrOp := (LIMITER - sftrOp)*(-1)
			nrc := rune(newSftrOp)
			nrx = append(nrx, nrc)
		}
	}

	return nrx
}

