package allergies

var allergens = [8]string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

func Allergies(score uint) []string {
	result := make([]string, 0)

	for _, allergen := range allergens {
		if AllergicTo(score, allergen) {
			result = append(result, allergen)
		}
	}

	return result
}

func AllergicTo(score uint, substance string) bool {
	ind := 0
	for score > 0 {
		if ind >= len(allergens) {
			break
		}

		if score&1 == 1 && allergens[ind] == substance {
			return true
		}

		score = score >> 1
		ind++
	}

	return false
}
