package pythagorean

type Triplet [3]int

func Range(min, max int) []Triplet {
	result := make([]Triplet, 0)

	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if a*a+b*b == c*c {
					result = append(result, Triplet{a, b, c})
				}
			}
		}
	}

	return result
}

func Sum(p int) []Triplet {
	result := make([]Triplet, 0)
	triplets := Range(1, p)
	for _, triplet := range triplets {
		if triplet[0]+triplet[1]+triplet[2] == p {
			result = append(result, triplet)
		}
	}
	return result
}
