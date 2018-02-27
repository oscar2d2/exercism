package house

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

func verseHelper(n, i int) string {
	if i == 1 {
		return "This is the " + phrase[n-i]
	}

	return verseHelper(n, i-1) + "\nthat " + verb[n-i] + " the " + phrase[n-i]
}

func Verse(n int) string {
	return verseHelper(n, n) + "."
}

func songHelper(i int) string {
	if i == 1 {
		return Verse(i)
	}

	return songHelper(i-1) + "\n\n" + Verse(i)
}

func Song() string {
	return songHelper(12)
}
