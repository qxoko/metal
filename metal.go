package metal

//go:generate go run util/gen.go util/source.txt

import (
	"time"
	"strings"
	"math/rand"
)

func Anodize(input string) string {
	// randomise the seed every time
	rand.Seed(time.Now().UnixNano())

	// we know characters will get wider, so
	// double the output buffer
	buffer := strings.Builder {}
	buffer.Grow(len(input) * 2)

	for _, c := range input {
		list, ok := charmap[c]

		if ok {
			n := len(list)

			if n > 1 {
				n -= 1
			}

			c = list[rand.Intn(n)]
		}

		buffer.WriteRune(c)
	}

	return buffer.String()
}