package scale

import (
	"strings"
)

var sharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flats = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

var sharpMap = map[string]bool{
	"a":  true,
	"C":  true,
	"G":  true,
	"D":  true,
	"A":  true,
	"E":  true,
	"B":  true,
	"F#": true,
	"e":  true,
	"b":  true,
	"f#": true,
	"c#": true,
	"g#": true,
	"d#": true,
}

var flatMap = map[string]bool{
	"F":  true,
	"Bb": true,
	"Eb": true,
	"Ab": true,
	"Db": true,
	"Gb": true,
	"d":  true,
	"g":  true,
	"c":  true,
	"f":  true,
	"bb": true,
	"eb": true,
}

func find(tonic string, list []string) int {
	tonic = strings.Title(tonic)

	for i, c := range list {
		if c == tonic {
			return i
		}
	}

	return -1
}

func Scale(tonic, interval string) []string {
	output := make([]string, 0)

	if _, ok := sharpMap[tonic]; ok {
		i := find(tonic, sharps)
		output = append(sharps[i:], sharps[0:i]...)
	}

	if _, ok := flatMap[tonic]; ok {
		i := find(tonic, flats)
		output = append(flats[i:], flats[0:i]...)
	}

	if interval != "" {
		temp := make([]string, len(interval))
		ind := 0
		temp[0] = output[ind]

		for i, step := range interval[0 : len(interval)-1] {
			if step == 'M' {
				ind += 2
			} else if step == 'm' {
				ind += 1
			} else if step == 'A' {
				ind += 3
			}
			temp[i+1] = output[ind]
		}
		output = temp
	}

	return output
}
