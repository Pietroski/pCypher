package check_max_size

import (
	"fmt"
	"github.com/Pietroski/pCypher/data/usecases/errors"
	"github.com/Pietroski/pCypher/data/usecases/messages"
	"github.com/Pietroski/timecounter"
	"github.com/Pietroski/timecounter/timescale"
	"testing"
)

type inputCMSParameters struct {
	initialRuneSliceSize int
	finalRuneSliceSize   int
	injectionFactor      int
}

type CMSStruct struct {
	input  inputCMSParameters
	output bool
}

var CheckMaxSizeMap = map[int]CMSStruct{
	0: {
		input:  inputCMSParameters{4, 14, 2},
		output: true,
	},
	1: {
		input:  inputCMSParameters{4, 9, 1},
		output: true,
	},
	2: {
		input:  inputCMSParameters{4, 9, 2},
		output: false,
	},
}

func CheckMaxSizeTest() {
	fmt.Println("CheckMaxSizeTest")
	for tn, t := range CheckMaxSizeMap {
		i := t.input
		o := t.output

		irs := i.initialRuneSliceSize
		frs := i.finalRuneSliceSize
		ijf := i.injectionFactor

		timecounter.TimeCounter(timescale.NANOSECONDS)
		timecounter.Start()
		fro := CheckMaxSize(irs, frs, ijf)
		timecounter.End()
		ps := messages.TestNumber(tn)
		timecounter.PrintTime(ps)

		if fro != o {
			errors.AssertionTestError(tn, o, fro)
		}
	}
	fmt.Println()
}

func TestCheckMaxSize(t *testing.T) {
	CheckMaxSizeTest()
}

func BenchmarkCheckMaxSize(b *testing.B) {
	CheckMaxSizeTest()
}

type inputMSParameters struct {
	runeSliceSize   int
	injectionFactor int
}

type MSStruct struct {
	input  inputMSParameters
	output int
}

var MaxSizeMap = map[int]MSStruct{
	0: {
		input:  inputMSParameters{4, 2},
		output: 14,
	},
}

func MaxSizeTest() {
	fmt.Println("MaxSizeTest")
	for tn, t := range MaxSizeMap {
		i := t.input
		o := t.output

		rs := i.runeSliceSize
		ijf := i.injectionFactor

		timecounter.TimeCounter(timescale.NANOSECONDS)
		timecounter.Start()
		fro := MaxSize(rs, ijf)
		timecounter.End()
		ps := messages.TestNumber(tn)
		timecounter.PrintTime(ps)

		if o != fro {
			errors.AssertionTestError(tn, o, fro)
		}
	}
	fmt.Println()
}

func TestMaxSize(t *testing.T) {
	MaxSizeTest()
}

func BenchmarkMaxSize(b *testing.B) {
	MaxSizeTest()
}
