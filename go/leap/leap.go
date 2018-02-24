// Check if a year is leap year or not
package leap

// IsLeapYear takes a int for year and return a bool to indicate if the year is a leap year
func IsLeapYear(year int) bool {

	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}

	return false
}
