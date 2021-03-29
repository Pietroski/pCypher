package inject_random_runes

import (
	"PietroskiCypher/data/usecases/errors"
	"PietroskiCypher/data/usecases/messages"
	"PietroskiCypher/data/usecontrollers/encode/check_max_size"

	"fmt"
	"github.com/Pietroski/timecounter/timescale"
	"log"
	"testing"
	"timecounter"
)

// InjectRandomRuneTest Tests
type IRRStruct struct {
	input           []rune
	output          []rune
	injectionFactor int
}

var InjectRandomRunesMap = map[int]IRRStruct{
	0: {
		input:           []rune{65, 66, 67, 68},
		output:          []rune{65, 87, 65, 54, 56, 66, 76, 34, 67, 98, 126, 68, 96, 67},
		injectionFactor: 2,
	},
	1: {
		input:           []rune{65, 66, 67, 68},
		output:          []rune{71, 36, 65, 71, 36, 66, 116, 105, 67, 63, 38, 68, 113, 86},
		injectionFactor: 2,
	},
}

func InjectRandomRunesTest() {
	fmt.Println("InjectRandomRuneTest")

	for tn, t := range InjectRandomRunesMap {
		i := t.input
		o := t.output
		ijf := t.injectionFactor

		timecounter.TimeCounter(timescale.NANOSECONDS)
		timecounter.Start()
		fro := InjectRandomRunes(i, ijf)
		timecounter.End()
		ps := messages.TestNumber(tn)
		timecounter.PrintTime(ps)

		iSeize := len(i)
		oSize := len(o)

		assrt := check_max_size.CheckMaxSize(iSeize, oSize, ijf)
		if !assrt {
			errors.AssertionTestError(tn, o, fro)
		}

		for indx := ijf; indx < len(o); {
			if indx < len(fro) && o[indx] != fro[indx] {
				errors.AssertionTestError(tn, o, fro)
			}
			indx += ijf + 1
		}

	}
	fmt.Println()
}

func TestInjectRandomRunes(t *testing.T) {
	InjectRandomRunesTest()
}

func BenchmarkInjectRandomRunes(b *testing.B) {
	InjectRandomRunesTest()
}


// RandomRuneSliceGenerator Tests
type RRSGStruct struct {
	input  int
	assertionFirst []rune
	assertionSecond int
}

var randomRuneSliceGeneratorMap = map[int]RRSGStruct{
	0: {
		input:          2,
		assertionFirst: []rune{65, 66},
		assertionSecond: 2,
	},
}

func RandomRuneSliceGeneratorTest() {
	fmt.Println("RandomRuneSliceGeneratorTest")
	for tn, t := range randomRuneSliceGeneratorMap {
		i := t.input
		a2 := t.assertionSecond

		timecounter.TimeCounter(timescale.NANOSECONDS)
		timecounter.Start()
		var fro interface{} = RandomRuneSliceGenerator(i)
		timecounter.End()
		ps := messages.TestNumber(tn)
		timecounter.PrintTime(ps)

		if fros, ok := fro.([]rune); !ok || len(fros) != a2 {
			terr := fmt.Errorf("type and/or length error on result(s) of test number -> [%v]", tn)
			log.Fatalln(terr)
		}
	}
	fmt.Println()
}

func TestRandomRuneSliceGenerator(t *testing.T) {
	RandomRuneSliceGeneratorTest()
}

func BenchmarkRandomRuneSliceGenerator(b *testing.B) {
	RandomRuneSliceGeneratorTest()
}


// RandomRuneGenerator Tests
type MinMaxStruct struct {
	min int
	max int
}

type RRGStruct struct {
	input  MinMaxStruct
	output rune
}

var RandomRuneGeneratorMap = map[int]RRGStruct{
	0: {
		input: MinMaxStruct{
			min: 33,
			max: 127,
		},
		output: rune(45),
	},
}

func RandomRuneGeneratorTest() {
	fmt.Println("RandomRuneGeneratorTest")
	for tn, t := range RandomRuneGeneratorMap {
		i := t.input

		timecounter.TimeCounter(timescale.NANOSECONDS)
		timecounter.Start()
		var fro interface{} = RandomRuneGenerator(i.min, i.max)
		timecounter.End()
		ps := messages.TestNumber(tn)
		timecounter.PrintTime(ps)

		if _, ok := fro.(rune); !ok {
			terr := fmt.Errorf("type error on result of test number -> [%v]", tn)
			log.Fatalln(terr)
		}
	}
	fmt.Println()
}

func TestRandomRuneGenerator(t *testing.T) {
	RandomRuneGeneratorTest()
}

func BenchmarkRandomRuneGenerator(b *testing.B) {
	RandomRuneGeneratorTest()
}
