package inject_random_runes

import (
	"PietroskiCypher/data/usecontrollers/encode/check_max_size"
	"log"
	"math/rand"
	"time"
)

// InjectRandomRunes into a slice of runes
func InjectRandomRunes(rs []rune, ijf int) []rune {
	rand.Seed(time.Now().UnixNano())

	rsSize := len(rs)

	var tmps []rune

	for i := 0; i < rsSize; i++ {
		tsr := RandomRuneSliceGenerator(ijf)

		if i == 0 {
			tmps = append(tsr)
		}

		srs := []rune{rs[i]}
		itmps := append(srs, tsr...)
		tmps = append(tmps, itmps...)
	}

	tmpsSize := len(tmps)
	vrf := check_max_size.CheckMaxSize(rsSize, tmpsSize, ijf)
	if !vrf {
		log.Fatalln("Error encrypting rune slice!!")
	}

	return tmps
}


// RandomRuneSliceGenerator generates a rune slice randomly of a fixed size
func RandomRuneSliceGenerator(ijf int) []rune {
	r := make([]rune, ijf)

	for i, _ := range r {
		r[i] = RandomRuneGenerator(33, 127)
	}

	return r
}


// RandomRuneGenerator generates a random rune between a
// minimum (including) and maximum (excluding) integer value
func RandomRuneGenerator(min int, max int) rune {
	rg := max - min
	rn := min + rand.Intn(rg)
	r := rune(rn)

	return r
}
