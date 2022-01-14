package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"github.com/qxoko/metal"
)

func main() {
	args := os.Args[1:]
	text := ""

	// if an argument exists, treat it as the
	// input.  otherwise, read the Stdin pipe
	if len(args) > 0 {
		text = args[0]
	} else {
		raw_text, err := ioutil.ReadAll(os.Stdin)

		if err != nil {
			panic(err)
		}

		text = string(raw_text)
	}

	// write output directly
	fmt.Fprintf(os.Stdout, metal.Anodize(text))

	// adds a newline for terminal output
	o, _ := os.Stdout.Stat()

	if (o.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		fmt.Println()
	}
}