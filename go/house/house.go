package house

import "bytes"

var phrase = []string{
	"house that Jack built",
	"malt",
	"rat",
	"cat",
	"dog",
	"cow with the crumpled horn",
	"maiden all forlorn",
	"man all tattered and torn",
	"priest all shaven and shorn",
	"rooster that crowed in the morn",
	"farmer sowing his corn",
	"horse and the hound and the horn",
}

var verb = []string{
	"lay in",
	"ate",
	"killed",
	"worried",
	"tossed",
	"milked",
	"kissed",
	"married",
	"woke",
	"kept",
	"belonged to",
}

func Verse(n int) string {
	var buffer bytes.Buffer

	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			buffer.WriteString("This is the " + phrase[i])
		} else {
			buffer.WriteString("\nthat " + verb[i] + " the " + phrase[i])
		}
	}

	buffer.WriteString(".")

	return buffer.String()
}

func Song() string {
	var buffer bytes.Buffer

	for i := 1; i <= 12; i++ {
		if i > 1 {
			buffer.WriteString("\n\n")
		}
		buffer.WriteString(Verse(i))
	}

	return buffer.String()
}
