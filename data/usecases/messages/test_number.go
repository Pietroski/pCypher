package messages

import "fmt"

// TestNumber returns a formatted string of the current
// test number passed as parameter
func TestNumber(tn int) string {
	ps := fmt.Sprintf("Test number %v took", tn)
	return ps
}
