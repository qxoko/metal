package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"github.com/qxoko/metal"
)

func main() {
	args := os.Args[1:]
	enum := metal.Vowels

	if len(args) > 0 {
		switch args[0] {
		case "--vowels":
			enum = metal.Vowels
		case "--all":
			enum = metal.All
		}
	}

	// read the catted file in
	raw_text, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	text := string(raw_text)

	// write output directly
	fmt.Fprintf(os.Stdout, metal.Anodize(text, enum))

	// this is supposed to detect a newline for output but we're
	// not quite at isatty levels of accuracy; basically it
	// doesn't work sometimes
	o, _ := os.Stdout.Stat()

	if (o.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		fmt.Println()
	}
}