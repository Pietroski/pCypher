package rune_slice_to_string

import (
	"PietroskiCypher/data/usecases/errors"
	"PietroskiCypher/data/usecases/messages"
	"fmt"
	"github.com/Pietroski/timecounter/timescale"
	"testing"
	"timecounter"
)

type RSSStruct struct {
	input  []rune
	output string
}

var RuneSliceToStringMap = map[int]RSSStruct{
	0: {
		input:  []rune{65, 83, 68, 70, 32, 97, 115, 100, 102},
		output: "ASDF asdf",
	},
	1: {
		input: []rune{
			65, 65, 65, 65, 65, 65, 65, 65, 65, 65,
			65, 65, 65, 65, 65, 65, 65, 65, 65, 65,
			65, 65, 65, 65, 65, 65, 65, 65, 65, 65,
			65, 65, 65, 65, 65, 65, 65, 65, 65, 65,
		},
		output: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
	},
	2: {
		input: []rune{65, 66, 66, 69, 70},
		output: "ABBEF",
	},
}

func RuneSliceToStringTest() {
	fmt.Println("RuneSliceToStringTest")

	for tn, ts := range RuneSliceToStringMap {
		i := ts.input
		o := ts.output

		timecounter.TimeCounter(timescale.NANOSECONDS)
		timecounter.Start()
		slts := RuneSliceToString(i)
		timecounter.End()
		ps := messages.TestNumber(tn)
		timecounter.PrintTime(ps)

		if slts != o {
			errors.AssertionTestError(tn, o, slts)
		}
	}
	fmt.Println()
}

func TestRuneSliceToString(t *testing.T) {
	RuneSliceToStringTest()
}

func BenchmarkRuneSliceToString(b *testing.B) {
	RuneSliceToStringTest()
}
