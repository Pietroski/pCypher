package string_to_rune_slice

import (
	"fmt"
	"github.com/Pietroski/pCypher/data/usecases/errors"
	"github.com/Pietroski/pCypher/data/usecases/messages"
	"github.com/Pietroski/timecounter"
	"github.com/Pietroski/timecounter/timescale"
	"reflect"
	"testing"
)

type STRSStruct struct {
	input  string
	output []rune
}

var StringToRuneSliceMap = map[int]STRSStruct{
	0: {
		input: "ASDF asdf",
		output: []rune{65, 83, 68, 70, 32, 97, 115, 100, 102},
	},
}

func StringToRuneSliceTest() {
	fmt.Println("StringToRuneSliceTest")

	for tn, ts := range StringToRuneSliceMap {
		i := ts.input
		o := ts.output

		timecounter.TimeCounter(timescale.NANOSECONDS)
		timecounter.Start()
		stobs := StringToRuneSlice(i)
		timecounter.End()
		ps := messages.TestNumber(tn)
		timecounter.PrintTime(ps)

		if !reflect.DeepEqual(stobs, o) {
			errors.AssertionTestError(tn, o, stobs)
		}
	}
	fmt.Println()
}

func TestStringToRuneSlice(t *testing.T) {
	StringToRuneSliceTest()
}

func BenchmarkStringToRuneSlice(b *testing.B) {
	StringToRuneSliceTest()
}
