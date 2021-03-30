package generics

import (
	"fmt"
	"github.com/Pietroski/timecounter"
	"github.com/Pietroski/timecounter/timescale"
	"reflect"
)

// Tester is used to execute and time function execution
func Tester(function interface{}, parameters ...interface{}) []reflect.Value {
	funcValue := reflect.ValueOf(function)
	funcParams := reflect.ValueOf(parameters)

	funcValueType := funcValue.Type()
	funcParamsType := funcParams.Type()

	if funcValueType.Kind() != reflect.Func {
		panic(fmt.Sprintf("function parameter is not of function type"))
	}

	if funcParamsType.Kind() != reflect.Slice && funcParamsType.Kind() != reflect.Array {
		panic("parameters parameter's type is neither array nor slice.")
	}

	funcParamsSize := funcParams.Len()

	var paramsSlice = make([]reflect.Value, 0, funcParamsSize)

	for i := 0; i < funcParamsSize; i++ {
		argsValues := funcParams.Index(i)
		argsValuesTypes := argsValues.Elem().Type()

		if !argsValuesTypes.ConvertibleTo(funcValueType.In(i)) {
			panic(fmt.Sprintf("Argument is not compatible with the function's argument number %v", i+1))
		}

		paramsSlice = append(paramsSlice, argsValues.Elem())
	}

	timecounter.TimeCounter(timescale.NANOSECONDS)
	timecounter.Start()
	fro := funcValue.Call(paramsSlice)
	timecounter.End()
	timecounter.PrintTime()

	return fro
}
