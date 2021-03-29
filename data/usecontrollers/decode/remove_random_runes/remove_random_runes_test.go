package remove_random_runes

import (
	"PietroskiCypher/data/usecases/errors"
	"PietroskiCypher/data/usecases/messages"
	"PietroskiCypher/data/usecontrollers/decode/check_min_size"
	"PietroskiCypher/data/usecontrollers/generics"

	"fmt"
	"github.com/Pietroski/timecounter/timescale"
	"testing"
	"timecounter"
)

// InjectRandomRuneTest Tests
type RRRStruct struct {
	input           []rune
	output          []rune
	injectionFactor int
}

var RemoveRandomRunesMap = map[int]RRRStruct{
	0: {
		input:           []rune{65, 87, 65, 54, 56, 66, 76, 34, 67, 98, 126, 68, 96, 67},
		output:          []rune{65, 66, 67, 68},
		injectionFactor: 2,
	},
}

var RemoveRandomRunesSlice = generics.MapToSlice(RemoveRandomRunesMap)

func RemoveRandomRuneSTest() {
	fmt.Println("RemoveRandomRuneSTest")
	RRRSSize := RemoveRandomRunesSlice.Len()
	for indx := 0; indx < RRRSSize; indx++ {
		v := RemoveRandomRunesSlice.Index(indx)

		pi := v.FieldByName("input")
		po := v.FieldByName("output")
		i := generics.SliceOfValuesIntoSliceOfRunes(pi)
		o := generics.SliceOfValuesIntoSliceOfRunes(po)
		injf := int(v.FieldByName("injectionFactor").Int())

		timecounter.TimeCounter(timescale.NANOSECONDS)
		timecounter.Start()
		fro := RemoveRandomRunes(i, injf)
		timecounter.End()
		ps := messages.TestNumber(indx)
		timecounter.PrintTime(ps)

		iSeize := len(i)
		oSize := len(o)

		assrt := check_min_size.CheckMinSize(iSeize, oSize, injf)
		if !assrt {
			errors.AssertionTestError(indx, o, fro)
		}

		for oindx := injf; oindx < len(o); {
			if oindx < len(fro) && o[oindx] != fro[oindx] {
				errors.AssertionTestError(indx, o, fro)
			}
			oindx += injf + 1
		}
	}
	fmt.Println()
}

func TestRemoveRandomRunes(t *testing.T) {
	RemoveRandomRuneSTest()
}

func BenchmarkRemoveRandomRunes(b *testing.B) {
	RemoveRandomRuneSTest()
}
