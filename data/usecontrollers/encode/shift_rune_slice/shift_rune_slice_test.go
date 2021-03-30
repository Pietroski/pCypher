package shift_rune_slice

import (
	"fmt"
	"github.com/Pietroski/pCypher/data/usecases/errors"
	"github.com/Pietroski/pCypher/data/usecases/messages"
	"github.com/Pietroski/pCypher/data/usecontrollers/generics"
	"github.com/Pietroski/timecounter"
	"github.com/Pietroski/timecounter/timescale"
	"reflect"
	"testing"
)

type SRSStruct struct {
	input   []rune
	output  []rune
	shifter int
}

var ShiftRuneSliceMap = map[int]SRSStruct{
	0: {
		input:   []rune{65, 83, 68, 70, 32, 97, 115, 100, 102},
		output:  []rune{78, 96, 81, 83, 45, 110, 128, 113, 115},
		shifter: 13,
	},
	1: {
		input:   []rune{1_112_064, 1_112_063, 1_112_062},
		output:  []rune{13, 12, 11},
		shifter: 13,
	},
}

func ShiftRuneSliceTest() {

	fmt.Println("ShiftRuneSliceTest")
	for tn, ts := range ShiftRuneSliceMap {
		i := ts.input
		o := ts.output
		s := ts.shifter

		timecounter.TimeCounter(timescale.NANOSECONDS)
		timecounter.Start()
		srs := ShiftRuneSlice(i, s)
		timecounter.End()
		ps := messages.TestNumber(tn)
		timecounter.PrintTime(ps)

		if !reflect.DeepEqual(srs, o) {
			errors.AssertionTestError(tn, o, srs)
		}
	}
	fmt.Println()

	fmt.Println("ShiftRuneSliceTest")
	for tn, ts := range ShiftRuneSliceMap {
		i := ts.input
		o := ts.output
		s := ts.shifter

		fro := generics.Tester(ShiftRuneSlice, i, s)[0]
		frox := generics.SliceOfValuesIntoSliceOfRunes(fro)

		if !reflect.DeepEqual(frox, o) {
			errors.AssertionTestError(tn, o, frox)
		}
	}
	fmt.Println()
}

func TestShiftRuneSlice(t *testing.T) {
	ShiftRuneSliceTest()
}

func BenchmarkShiftRuneSlice(b *testing.B) {
	ShiftRuneSliceTest()
}