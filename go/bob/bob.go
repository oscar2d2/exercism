package bob

import (
	"regexp"
	"strings"
)

func Hey(remark string) string {
	var reply string
	r, _ := regexp.Compile(`^[\d,\s?]+$`)
	r2, _ := regexp.Compile(`[a-zA-Z]+`)

	remark = strings.Trim(remark, " \t\r\n")

	switch {
	case remark == "":
		reply = "Fine. Be that way!"
	case r.MatchString(remark) && strings.ContainsAny(remark, "?"):
		reply = "Sure."
	case strings.Compare(strings.ToUpper(remark), remark) == 0 && strings.ContainsAny(remark, "?") && r2.MatchString(remark):
		reply = "Calm down, I know what I'm doing!"
	case r.MatchString(remark):
		reply = "Whatever."
	case strings.Compare(strings.ToUpper(remark), remark) == 0 && r2.MatchString(remark):
		reply = "Whoa, chill out!"
	case strings.HasSuffix(remark, "?"):
		reply = "Sure."
	default:
		reply = "Whatever."
	}

	return reply
}
