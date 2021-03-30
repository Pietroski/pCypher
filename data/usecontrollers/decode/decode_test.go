package decode

import (
	"fmt"
	"github.com/Pietroski/pCypher/data/usecases/errors"
	"github.com/Pietroski/pCypher/data/usecases/messages"
	"github.com/Pietroski/pCypher/data/usecontrollers/decode/check_min_size"
	"github.com/Pietroski/pCypher/data/usecontrollers/generics"
	"github.com/Pietroski/timecounter"
	"github.com/Pietroski/timecounter/timescale"
	"testing"
)

// DStruct struct type
type DStruct struct {
	input           string
	output          string
	shifter         int
	injectionFactor int
}

var DecodeMap = map[int]DStruct{
	0: {
		input:           "vNdNgNwNgNa",
		output:          "AAAAA",
		shifter:         13,
		injectionFactor: 1,
	},
	1: {
		input:           "NNNNN",
		output:          "AAAAA",
		shifter:         13,
		injectionFactor: 0,
	},
	2: {
		input:           "aNaNaNaNaNa",
		output:          "AAAAA",
		shifter:         13,
		injectionFactor: 1,
	},
	3: {
		input:           "aNaNaiaNaNaNa",
		output:          "AA\\AAA",
		shifter:         13,
		injectionFactor: 1,
	},
}

var DecodeSlice = generics.MapToSlice(DecodeMap)

func DecodeTest() {
	fmt.Println("DecodeTest")

	for indx := 0; indx < DecodeSlice.Len(); indx++ {
		v := DecodeSlice.Index(indx)

		i := v.FieldByName("input").String()
		o := v.FieldByName("output").String()
		sftr := int(v.FieldByName("shifter").Int())
		injf := int(v.FieldByName("injectionFactor").Int())

		timecounter.TimeCounter(timescale.NANOSECONDS)
		timecounter.Start()
		fro := Decode(i, sftr, injf)
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

	fmt.Println("DecodeTest 2")
	for tn, ts := range DecodeMap {
		i := ts.input
		o := ts.output
		s := ts.shifter
		ijf := ts.injectionFactor

		timecounter.TimeCounter(timescale.NANOSECONDS)
		timecounter.Start()
		fro := Decode(i, s, ijf)
		timecounter.End()
		ps := messages.TestNumber(tn)
		timecounter.PrintTime(ps)


		iSeize := len(i)
		oSize := len(o)

		assrt := check_min_size.CheckMinSize(iSeize, oSize, ijf)
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

func TestDecode(t *testing.T) {
	DecodeTest()
}

func BenchmarkDecode(b *testing.B) {
	DecodeTest()
}
