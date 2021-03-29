package unshift_rune_slice

import (
	"PietroskiCypher/data/usecases/errors"
	"PietroskiCypher/data/usecases/messages"
	"fmt"
	"github.com/Pietroski/timecounter/timescale"
	"reflect"
	"testing"
	"timecounter"
)

type USRSStruct struct {
	input  []rune
	output []rune
	shifter int
}

var UnshiftRuneSliceMap = map[int]USRSStruct{
	0: {
		input:  []rune{78, 96, 81, 83, 45, 110, 128, 113, 115},
		output: []rune{65, 83, 68, 70, 32, 97, 115, 100, 102},
		shifter: 13,
	},
	1: {
		input:  []rune{13, 12, 11},
		output: []rune{0, 1_112_063, 1_112_062},
		shifter: 13,
	},
}

func UnshiftRuneSliceTest() {

	fmt.Println("UnshiftRuneSliceTest")
	for tn, ts := range UnshiftRuneSliceMap {
		i := ts.input
		o := ts.output
		s := ts.shifter

		timecounter.TimeCounter(timescale.NANOSECONDS)
		timecounter.Start()
		srs := UnshiftRuneSlice(i, s)
		timecounter.End()

		ps := messages.TestNumber(tn)
		timecounter.PrintTime(ps)

		if !reflect.DeepEqual(srs, o) {
			errors.AssertionTestError(tn, o, srs)
		}
	}
	fmt.Println()
}

func TestUnshiftRuneSliceTest(t *testing.T) {
	UnshiftRuneSliceTest()
}

func BenchmarkUnshiftRuneSliceTest(b *testing.B) {
	UnshiftRuneSliceTest()
}
