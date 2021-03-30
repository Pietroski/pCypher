package decode

import (
	"github.com/Pietroski/pCypher/data/usecontrollers/decode/remove_random_runes"
	"github.com/Pietroski/pCypher/data/usecontrollers/decode/unshift_rune_slice"
	"github.com/Pietroski/pCypher/data/usecontrollers/generics/rune_slice_to_string"
	"github.com/Pietroski/pCypher/data/usecontrollers/generics/string_to_rune_slice"
)

// Decode encodes any string
func Decode(stringToDecode string, shiftFactor int, injectionFactor int) string {
	strx := string_to_rune_slice.StringToRuneSlice(stringToDecode)
	srx := unshift_rune_slice.UnshiftRuneSlice(strx, shiftFactor)
	rrr := remove_random_runes.RemoveRandomRunes(srx, injectionFactor)
	rxts := rune_slice_to_string.RuneSliceToString(rrr)
	return rxts
}
