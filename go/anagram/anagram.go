package anagram

import (
	"sort"
	"strings"
)

type MyString []rune

func (s MyString) Len() int {
	return len(s)
}
func (s MyString) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s MyString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func isCapital(s string) bool {
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			return false
		}
	}

	return true
}

func Detect(subject string, candidates []string) []string {
	isCap := isCapital(subject)
	result := make([]string, 0)
	a := MyString(strings.ToLower(subject))
	sort.Sort(a)

	for _, s := range candidates {
		if isCap && strings.ToLower(subject) == strings.ToLower(s) {
			continue
		}
		temp := MyString(strings.ToLower(s))
		sort.Sort(temp)
		if string(temp) == string(a) {
			result = append(result, s)
		}
	}
	return result
}
