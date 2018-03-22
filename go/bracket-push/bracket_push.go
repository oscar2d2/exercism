package brackets

var m = map[rune]rune{
	']': '[',
	')': '(',
	'}': '{',
}

func Bracket(input string) (result bool, err error) {
	s := make([]rune, 0)
	result = false

	for _, c := range input {
		if c == ']' || c == '}' || c == ')' {
			if len(s) == 0 || s[len(s)-1] != m[c] {
				return
			} else {
				s = s[:len(s)-1]
			}
		} else if c == '[' || c == '{' || c == '(' {
			s = append(s, c)
		}
	}

	if len(s) == 0 {
		result = true
	}

	return
}
