# ɱéťàł

> *"English, but metal"*

Metal is a tool that converts English text a legible, Zalgo-like character swap for the purposes of testing localisation and UTF-8 character support in game engines, websites and other tools.

By rewriting basic ASCII characters into recognisable, accented alternatives, an English-reading programmer can immediately see where and how text is breaking in a rendering pipeline.

It should be clear, however, that Metal is not a magic bullet for all text rendering or all language support, merely an early testing tool.

#### Before

> Alice was beginning to get very tired of sitting by her sister on the bank, and of having nothing to do:  once or twice she had peeped into the book her sister was reading, but it had no pictures or conversations in it, `and what is the use of a book,' thought Alice `without pictures or conversation?'

#### After

> Ăłɨĉə ŵăŝ ьèğîɴɲĭŉġ ţö ɠêţ vɘɹý ʈìʀəđ ôf šíťʈĩɴġ ʙŷ ɦěʀ šĩşţèŗ óŋ ŧĥę ьàɲʞ, åňɗ őf ɧąvîñğ ņöťɦĩŉɠ ţö ɗŏ:  øɴčê ŏʀ ʈŵĭčę
šɧê ħăɗ pěəpĕɗ ìñţô ťɧĕ ьøõk ħěɾ ŝīŝťėŕ ŵąŝ ŗêàɗĭńġ, ʙûţ íʈ ɦāɗ ŉő pîćʈůʀéş ōř çôŋvəřŝâťĭøňś īń ìŧ, ŵïţĥŏúť pĩçţuřęş őʁ
çõŋvęɹśáŧïòŋ?'

Metal was inspired by a segment on internationalisation in [this GDC talk about *Firewatch*](https://www.youtube.com/watch?v=wj-2vbiyHnI), given by William Armstrong and Patrick Ewing.

## Standalone Usage

	$ cat input.txt | metal > output.txt

	$ metal "some input string"

## Standalone Compilation

	$ go build -o metal<.exe> exec/main.go

The character table can also be modified or rebuilt from the `source.txt` file in `util`.  Run —

	$ go generate

— to rebuild it.

## Library Usage

```go
import "github.com/qxoko/metal"

some_output := metal.Anodize(some_input)
```