package secret

var action = map[uint]string{
	1: "wink",
	2: "double blink",
	4: "close your eyes",
	8: "jump",
}

func reverse(l []string) []string {
	var out = make([]string, len(l))

	for i := 0; i < len(l); i++ {
		out[i] = l[len(l)-1-i]
	}

	return out
}

func Handshake(code uint) []string {
	var output = make([]string, 0)
	var bitmask uint = 8

	for bitmask > 0 {
		if code&bitmask > 0 {
			output = append(output, action[bitmask])
		}
		bitmask = bitmask >> 1
	}

	if code < 16 {
		output = reverse(output)
	}

	return output
}
