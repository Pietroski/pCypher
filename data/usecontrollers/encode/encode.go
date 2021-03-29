package encode

import (
	"PietroskiCypher/data/usecontrollers/encode/inject_random_runes"
	"PietroskiCypher/data/usecontrollers/encode/shift_rune_slice"
	"PietroskiCypher/data/usecontrollers/generics/rune_slice_to_string"
	"PietroskiCypher/data/usecontrollers/generics/string_to_rune_slice"
)

// Encode encodes any string
func Encode(stringToEncode string, shiftFactor int, injectionFactor int) string {
	strx := string_to_rune_slice.StringToRuneSlice(stringToEncode)
	srx := shift_rune_slice.ShiftRuneSlice(strx, shiftFactor)
	irr := inject_random_runes.InjectRandomRunes(srx, injectionFactor)
	rxts := rune_slice_to_string.RuneSliceToString(irr)
	return rxts
}
