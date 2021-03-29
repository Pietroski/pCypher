package pCypher

// PCypher interface
type PCypher interface {
	Encode(stringToEncode string, shiftFactor int, injectionFactor int) string
	Decode(stringToDecode string, shiftFactor int, injectionFactor int) string
}
