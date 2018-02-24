// Package leap if a year is leap year or not
package leap

// IsLeapYear takes a int for year and return a bool to indicate if the year is a leap year
func IsLeapYear(year int) bool {
	return year%400 == 0 || (year%100 != 0 && year%4 == 0)
}
