package generics

import (
	"fmt"
	"testing"
)

func HelloTest(s string) []string {
	if s == "" {
		fmt.Println("NOTHING TO SHOW HERE")
	}

	fmt.Println("JUST BEFORE RETURN")
	fmt.Println(s)
	return []string{s, "another", "test"}
}

func TestTester(t *testing.T) {
	fmt.Println("TestTester")

	fro := Tester(HelloTest, "Hello My Test")[0]

	fmt.Println("FRO ->", fro)
}
