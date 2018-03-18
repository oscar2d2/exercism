package palindrome

import (
	"errors"
	"fmt"
)

type Product struct {
	Product        int
	Factorizations [][2]int
}

func isPalidrome(x int) bool {
	return isPalidromeHelper(fmt.Sprintf("%v", x))
}

func isPalidromeHelper(s string) bool {
	if len(s) <= 1 {
		return true
	}

	if s[0] == s[len(s)-1] {
		return isPalidromeHelper(s[1 : len(s)-1])
	}

	return false
}

func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		err = errors.New("fmin > fmax...")
		return
	}

	pmin = Product{Product: fmin * fmax}
	pmax = Product{Product: 0}

	for first := fmin; first <= fmax; first++ {
		for second := first; second <= fmax; second++ {
			product := first * second
			if !isPalidrome(product) {
				continue
			}
			if product < pmin.Product {
				pmin.Product = product
				pmin.Factorizations = append([][2]int{}, [2]int{first, second})
			} else if product == pmin.Product {
				pmin.Factorizations = append(pmin.Factorizations, [2]int{first, second})
			}

			if product > pmax.Product {
				pmax.Product = product
				pmax.Factorizations = append([][2]int{}, [2]int{first, second})
			} else if product == pmax.Product {
				pmax.Factorizations = append(pmax.Factorizations, [2]int{first, second})
			}
		}
	}

	if pmax.Product == 0 {
		err = errors.New("no palindromes...")
		return
	}

	return
}
