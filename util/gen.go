package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"io/ioutil"
	"path/filepath"
)

/*
	all .txt files will be converted into maps
	see existing .txt files for formatting

	map names are based on their filenames:
	<name>.txt  -->  <name>map = map[]...

	from the root directory:
	go run util/gen.go util <target_name.go>
*/

const boilerplate         = "package metal\n\n"
const map_entry_template  = "var %s = map[rune][]rune {\n\t%s\n}\n\n"
const line_entry_template = "\t'%s': []rune{%s},"

func main() {
	// this is all horrid but it's not hot path
	// so i'm not massively upset by how garbage
	// all these different file access methods are

	util_path := os.Args[1:][0]

	fs, err := ioutil.ReadDir(util_path)

	if err != nil {
		panic(err)
	}

	buffer := strings.Builder {}

	buffer.Grow(len(boilerplate))
	buffer.WriteString(boilerplate)

	for _, f := range fs {
		name := filepath.Join(util_path, f.Name())

		if strings.HasSuffix(name, ".txt") {
			text := do_file(name)

			buffer.Grow(len(text))
			buffer.WriteString(text)

			fmt.Println(f.Name())
		}
	}

	ioutil.WriteFile(os.Args[1:][1], []byte(strings.TrimSpace(buffer.String())), 0777)
}

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

func do_file(file_name string) string {
	file, err := os.Open(file_name)

	defer file.Close()

	if err != nil {
		panic(err)
	}

	info, err := file.Stat()

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

	map_name := filepath.Base(file_name)
	map_name  = map_name[:len(map_name)-4]

	return fmt.Sprintf(map_entry_template, map_name + "map", strings.TrimSpace(buffer.String()))
}