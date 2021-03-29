package check_min_size

import (
	"PietroskiCypher/data/usecases/errors"
	"PietroskiCypher/data/usecases/messages"
	"fmt"
	"github.com/Pietroski/timecounter/timescale"
	"testing"
	"timecounter"
)

type inputCMinSParameters struct {
	initialRuneSliceSize int
	finalRuneSliceSize   int
	injectionFactor      int
}

type CMinSStruct struct {
	input  inputCMinSParameters
	output bool
}

var CheckMinSizeMap = map[int]CMinSStruct{
	0: {
		input:  inputCMinSParameters{14, 4, 2},
		output: true,
	},
	1: {
		input:  inputCMinSParameters{9, 4, 1},
		output: true,
	},
	2: {
		input:  inputCMinSParameters{9, 4, 2},
		output: false,
	},
}

func CheckMinSizeTest() {
	fmt.Println("CheckMinSizeTest")
	for tn, t := range CheckMinSizeMap {
		i := t.input
		o := t.output

		irs := i.initialRuneSliceSize
		frs := i.finalRuneSliceSize
		injf := i.injectionFactor

		timecounter.TimeCounter(timescale.NANOSECONDS)
		timecounter.Start()
		fro := CheckMinSize(irs, frs, injf)
		timecounter.End()
		ps := messages.TestNumber(tn)
		timecounter.PrintTime(ps)

		if fro != o {
			errors.AssertionTestError(tn, o, fro)
		}
	}
	fmt.Println()
}

func TestCheckMinSize(t *testing.T) {
	CheckMinSizeTest()
}

func BenchmarkCheckMinSize(b *testing.B) {
	CheckMinSizeTest()
}

type inputMinSParameters struct {
	runeSliceSize   int
	injectionFactor int
}

type MinSStruct struct {
	input  inputMinSParameters
	output int
}

var MinSizeMap = map[int]MinSStruct{
	0: {
		input:  inputMinSParameters{14, 2},
		output: 4,
	},
	1: {
		input:  inputMinSParameters{9, 1},
		output: 4,
	},
}

func MinSizeTest() {
	fmt.Println("MinSizeTest")
	for tn, t := range MinSizeMap {
		i := t.input
		o := t.output

		rs := i.runeSliceSize
		injf := i.injectionFactor

		timecounter.TimeCounter(timescale.NANOSECONDS)
		timecounter.Start()
		fro := MinSize(rs, injf)
		timecounter.End()
		ps := messages.TestNumber(tn)
		timecounter.PrintTime(ps)

		if o != fro {
			errors.AssertionTestError(tn, o, fro)
		}
	}
	fmt.Println()
}

func TestMinSize(t *testing.T) {
	MinSizeTest()
}

func BenchmarkMinSize(b *testing.B) {
	MinSizeTest()
}
