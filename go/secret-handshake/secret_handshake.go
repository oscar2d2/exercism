package secret

var action = map[uint]string{
	1: "wink",
	2: "double blink",
	4: "close your eyes",
	8: "jump",
}

var actions = []uint{8, 4, 2, 1}

func Handshake(code uint) []string {
	var left2right bool
	var output = make([]string, 1)

	if code >= 16 {
		left2right = true
		code -= 16
	}

	for _, a := range actions {
		if code >= a {
			output = append(output, action[a])
			code -= a
		}
	}

	output = output[1:]

	if !left2right {
		for i := 0; i < len(output)/2; i++ {
			output[i], output[len(output)-1-i] = output[len(output)-1-i], output[i]
		}
	}

	return output
}
