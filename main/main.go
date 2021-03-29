package main

import (
	"PietroskiCypher/data/usecontrollers/decode"
	"PietroskiCypher/data/usecontrollers/encode"
	"fmt"
)

func main() {
	estr := encode.Encode("á'^~á ~aãâ\\", 13, 2)
	fmt.Println(estr)
	dstr := decode.Decode(estr, 13, 2)
	fmt.Println(dstr)
}
