package metal

//go:generate go run util/gen.go util chars.go

import (
	"time"
	"strings"
	"math/rand"
)

const (
	All uint8 = iota
	Vowels
)

func Anodize(input string, choice uint8) string {
	// randomise the seed every time
	rand.Seed(time.Now().UnixNano())

	if choice == Vowels {
		return do_substitution(input, vowelmap)
	}

	return do_substitution(input, charmap)
}

func do_substitution(input string, the_map map[rune][]rune) string {
	// we know characters will get wider, so
	// double the output buffer
	buffer := strings.Builder {}
	buffer.Grow(len(input) * 2)

	for _, c := range input {
		list, ok := the_map[c]

		if ok {
			n := len(list)
			if n > 1 { n -= 1 }
			c = list[rand.Intn(n)]
		}

		buffer.WriteRune(c)
	}

	return buffer.String()
}