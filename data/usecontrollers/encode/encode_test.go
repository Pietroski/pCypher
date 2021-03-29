package encode

import (
	"PietroskiCypher/data/usecases/errors"
	"PietroskiCypher/data/usecases/messages"
	"PietroskiCypher/data/usecontrollers/encode/check_max_size"
	"PietroskiCypher/data/usecontrollers/generics"
	"fmt"
	"github.com/Pietroski/timecounter/timescale"
	"testing"
	"timecounter"
)

// EStruct struct type
type EStruct struct {
	input           string
	output          string
	shifter         int
	injectionFactor int
}

var EncodeMap = map[int]EStruct{
	0: {
		input:           "AAAAA",
		output:          "vNdNgNwNgNa",
		shifter:         13,
		injectionFactor: 1,
	},
	1: {
		input:           "AAAAA",
		output:          "NNNNN",
		shifter:         13,
		injectionFactor: 0,
	},
	2: {
		input:           "AAAAA",
		output:          "aNaNaNaNaNa",
		shifter:         13,
		injectionFactor: 1,
	},
	3: {
		input:           "AAAAA",
		output:          "aNaNaNaNaNa",
		shifter:         13,
		injectionFactor: 1,
	},
	4: {
		input:           "AA\\AAA",
		output:          "aNaNaiaNaNaNa",
		shifter:         13,
		injectionFactor: 1,
	},
}

var EncodeSlice = generics.MapToSlice(EncodeMap)

func EncodeTest() {
	fmt.Println("EncodeTest")

	for indx := 0; indx < EncodeSlice.Len(); indx++ {
		v := EncodeSlice.Index(indx)

		i := v.FieldByName("input").String()
		o := v.FieldByName("output").String()
		sftr := int(v.FieldByName("shifter").Int())
		injf := int(v.FieldByName("injectionFactor").Int())

		timecounter.TimeCounter(timescale.NANOSECONDS)
		timecounter.Start()
		fro := Encode(i, sftr, injf)
		timecounter.End()
		ps := messages.TestNumber(indx)
		timecounter.PrintTime(ps)

		iSeize := len(i)
		oSize := len(o)

		assrt := check_max_size.CheckMaxSize(iSeize, oSize, injf)
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

	fmt.Println("EncodeTest 2")
	for tn, ts := range EncodeMap {
		i := ts.input
		o := ts.output
		s := ts.shifter
		ijf := ts.injectionFactor

		timecounter.TimeCounter(timescale.NANOSECONDS)
		timecounter.Start()
		fro := Encode(i, s, ijf)
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

func TestEncode(t *testing.T) {
	EncodeTest()
}

func BenchmarkEncode(b *testing.B) {
	EncodeTest()
}
