package sublist

type Relation string

func Sublist(l1, l2 []int) Relation {
	if len(l1) == 0 && len(l2) == 0 {
		return "equal"
	} else if len(l1) == 0 {
		return "sublist"
	} else if len(l2) == 0 {
		return "superlist"
	}

	var as, bs []int

	if len(l1) <= len(l2) {
		as, bs = l1, l2
	} else {
		as, bs = l2, l1
	}

	for i := 0; i < len(bs); i++ {
		for j := 0; j < len(as); j++ {
			if i+j >= len(bs) || as[j] != bs[i+j] {
				break
			} else if j == len(as)-1 {
				if len(l1) == len(l2) {
					return "equal"
				} else if len(l1) < len(l2) {
					return "sublist"
				} else {
					return "superlist"
				}
			}
		}
	}

	return "unequal"
}
