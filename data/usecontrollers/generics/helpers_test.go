package generics

import (
	"PietroskiCypher/data/usecases/errors"
	"fmt"
	"reflect"
	"testing"
)

type UniversalMap map[interface{}]interface{}
type UniversalSlice []interface{}

type MToSStruct struct {
	input  UniversalMap
	output UniversalSlice
}

var MtoSMap = map[int]MToSStruct{
	0: {
		input:  map[interface{}]interface{}{0: "first", 1: "test"},
		output: []interface{}{"first", "test"},
	},
}

func TestMapToSlice(t *testing.T) {
	fmt.Println("TestMapToSlice")
	for tn, t := range MtoSMap {
		i := t.input
		o := t.output

		slicedMap := MapToSlice(i)

		slicedMapSize := slicedMap.Len()

		for i := 0; i < slicedMapSize; i++ {
			sldMS := slicedMap.Index(i)
			os := o[i]

			if reflect.DeepEqual(sldMS, os) {
				errors.AssertionTestError(tn, o, slicedMap)
			}
		}
	}

}

func TestSliceOfValuesIntoSliceOfRunes(t *testing.T) {
	fmt.Println("TestSliceOfValuesIntoSliceOfRunes")

	i := []rune{65, 66, 67, 68}
	o := []rune{65, 66, 67, 68}

	SOfV := reflect.ValueOf(i)

	fmt.Println(SOfV.Len(), o, SOfV.Type(), SOfV)
	fro := SliceOfValuesIntoSliceOfRunes(SOfV)

	if !reflect.DeepEqual(fro, o) {
		errors.AssertionTestError(0, o, fro)
	}
	fmt.Println()
}
