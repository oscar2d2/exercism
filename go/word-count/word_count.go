package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := make(Frequency, 0)

	f := func(r rune) bool {
		return (unicode.IsPunct(r) && r != '\'') || unicode.IsSpace(r) || unicode.IsSymbol(r)
	}

	phrase = strings.ToLower(phrase)
	for _, s := range strings.FieldsFunc(phrase, f) {
		if s[0] == '\'' && s[len(s)-1] == '\'' {
			s = s[1 : len(s)-1]
		}
		freq[s]++
	}

	return freq
}
