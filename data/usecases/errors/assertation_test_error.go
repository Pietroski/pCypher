package errors

import (
	"fmt"
	"log"
)

// AssertionTestError Handles a fatal error for assertion mistake
// it receives as parameter the test number, the expected output and
// the function returned result, respectively
func AssertionTestError(tn int, i ...interface{}) {
	o := i[0]
	fro := i[1]

	fmt.Printf("\nAssertation failed on test number %v!!\n", tn)
	err := fmt.Errorf("\n Expected -> %v\n Returned -> %v\n", o, fro)
	log.Fatalln(err)
}
