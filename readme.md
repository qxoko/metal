# ɱéťàł

> *"English, but metal"*

Metal is a tool that converts English text into a legible, Zalgo-like character swap for the purposes of testing localisation and UTF-8 character support in game engines, websites and other tools.

By rewriting basic ASCII characters into recognisable, accented alternatives, an English-reading programmer can immediately see where and how text is breaking in a rendering pipeline.

It should be clear, however, that Metal is not a magic bullet for all text rendering or all language support, merely an early testing tool.

#### Before

> Alice was beginning to get very tired of sitting by her sister on the bank, and of having nothing to do:  once or twice she had peeped into the book her sister was reading, but it had no pictures or conversations in it, 'and what is the use of a book,' thought Alice 'without pictures or conversation?'

#### After

> Âlìcę wãs bègínnìng tó gēt vërý tïrëd óf sįttĩng bŷ hęr sĭstėr õn thė bānk, ãnd őf hăvįng nöthìng tő dö:  òncĕ õr twĭcé shë håd péêpëd ìntó thë bóők hēr sįstĕr wãs rĕådĭng, bŭt ït håd nõ píctürës ōr cŏnvěrsätįôns ïn ìt, 'ánd whãt ïs thē usę ɵf ă bòôk,' thóűght Älícé 'wĩthõut pïcturęs ór cônvèrsåtìŏn?'

Metal was inspired by a segment on internationalisation in [this GDC talk about *Firewatch*](https://www.youtube.com/watch?v=wj-2vbiyHnI), given by William Armstrong and Patrick Ewing.

## Library Usage

```go
import "github.com/qxoko/metal"

some_output := metal.Anodize(some_input, metal.Vowels)
```

The available enums are:

	metal.All
	metal.Vowels

## Standalone Usage

	$ cat input.txt | metal --all > output.txt

The optional flags are:

	--vowels (default)
	--all

...where "vowels" will only do vowel substitution, but "all" will provide consonants as well.

## Standalone Compilation

	$ go build -o metal<.exe> exec/main.go

The character table can also be modified or rebuilt from the `source.txt` file in `util`.  Run —

	$ go generate

— to rebuild it.