package twelve

import (
	"bytes"
	"fmt"
)

var phrase = map[int]string{
	1:  "a Partridge in a Pear Tree",
	2:  "two Turtle Doves",
	3:  "three French Hens",
	4:  "four Calling Birds",
	5:  "five Gold Rings",
	6:  "six Geese-a-Laying",
	7:  "seven Swans-a-Swimming",
	8:  "eight Maids-a-Milking",
	9:  "nine Ladies Dancing",
	10: "ten Lords-a-Leaping",
	11: "eleven Pipers Piping",
	12: "twelve Drummers Drumming",
}

var dayth = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

func Song() string {
	var buffer bytes.Buffer
	for i := 1; i <= 12; i++ {
		buffer.WriteString(Verse(i) + "\n")
	}
	return buffer.String()
}

func Verse(day int) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me,%s", dayth[day], helper(day, day))
}

func helper(day, n int) string {
	if day == 1 {
		if n > 1 {
			return " and " + phrase[day] + "."
		}
		return " " + phrase[day] + "."
	}

	return " " + phrase[day] + "," + helper(day-1, n)
}
