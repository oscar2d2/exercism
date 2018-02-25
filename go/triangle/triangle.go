// Package triangle checks the type of triangle
package triangle

import "math"

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides take lenght of three sides and return the kind of the triangle
func KindFromSides(a, b, c float64) Kind {
	if !validate(a, b, c) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	} else if a == b || b == c || a == c {
		return Iso
	} else {
		return Sca
	}

	return NaT
}

func validate(a, b, c float64) bool {

	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return false
	}

	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return false
	}

	if a+b < c || a+c < b || b+c < a {
		return false
	}

	return true
}
