package foodchain

import (
	"bytes"
	"fmt"
)

var animals = [8]string{"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}

var sentences = map[string]string{
	"fly":    "I don't know why she swallowed the fly. Perhaps she'll die.",
	"spider": "It wriggled and jiggled and tickled inside her.",
	"bird":   "How absurd to swallow a bird!",
	"cat":    "Imagine that, to swallow a cat!",
	"dog":    "What a hog, to swallow a dog!",
	"goat":   "Just opened her throat and swallowed a goat!",
	"cow":    "I don't know how she swallowed a cow!",
	"horse":  "She's dead, of course!",
}

func Verse(v int) string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("I know an old lady who swallowed a %s.\n", animals[v-1]))

	buffer.WriteString(sentences[animals[v-1]])

	if v == 1 || v == 8 {
		return buffer.String()
	}

	for i := v - 1; i >= 1; i-- {
		buffer.WriteString("\n")
		if v > 2 && i == 2 {
			buffer.WriteString(fmt.Sprintf("She swallowed the %s to catch the %s that %s", animals[i], animals[i-1], sentences[animals[i-1]][3:]))
		} else {
			buffer.WriteString(fmt.Sprintf("She swallowed the %s to catch the %s.", animals[i], animals[i-1]))
		}
	}
	buffer.WriteString("\n")
	buffer.WriteString(sentences[animals[0]])

	return buffer.String()
}

func Verses(a, b int) string {
	var buffer bytes.Buffer

	for i := a; i <= b; i++ {
		if i > a {
			buffer.WriteString("\n\n")
		}
		buffer.WriteString(Verse(i))
	}

	return buffer.String()
}

func Song() string {
	return Verses(1, 8)
}
