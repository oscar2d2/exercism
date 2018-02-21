package acronym

import (
	"bytes"
	"strings"
)

func Abbreviate(s string) string {
	var buffer bytes.Buffer
	strs := strings.FieldsFunc(s, func(r rune) bool {
		switch r {
		case ' ', '-':
			return true
		}

		return false
	})

	for _, str := range strs {
		buffer.WriteString(str[:1])
	}

	return strings.ToUpper(buffer.String())
}
