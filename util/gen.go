//go:build ignore

package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"io/ioutil"
)

const file_name   = `chars.go`
const boilerplate = `package metal

var charmap = map[rune][]rune {
	%s
}`

const line_entry_template = `	'%s': []rune{%s},`

func line_entry(input string) string {
	buffer := strings.Builder {}
	buffer.Grow(len(input) * 4)

	for _, c := range input {
		buffer.WriteRune('\'')
		buffer.WriteRune(c)
		buffer.WriteString("',")
	}

	bake := buffer.String()

	return bake[:len(bake) - 1] // remove trailing comma
}

func main() {
	file, err := os.Open(os.Args[1:][0])

	defer file.Close()

	if err != nil {
		panic(err)
	}

	info, err := file.Stat()

	if err != nil {
		panic(err)
	}

	buffer := strings.Builder {}
	buffer.Grow(int(info.Size()) * 2)

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) < 2 {
			continue
		}

		key_lo := string(line[0])
		key_hi := strings.ToUpper(key_lo)

		lo := line[2:]
		hi := strings.ToUpper(lo)

		buffer.WriteString(fmt.Sprintf(line_entry_template, key_lo, line_entry(lo)))
		buffer.WriteRune('\n')

		buffer.WriteString(fmt.Sprintf(line_entry_template, key_hi, line_entry(hi)))
		buffer.WriteRune('\n')
	}

	out := strings.TrimSpace(buffer.String())

	ioutil.WriteFile(file_name, []byte(fmt.Sprintf(boilerplate, out)), 0777)
}