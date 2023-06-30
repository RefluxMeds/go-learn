// Package dog contains a subset of functions which allows users to operate on dog objects
package dogs

// ToDogYears function will take in human years as a parameter and return the corresponding dog years back
func ToDogYears(hy int) int {
	dy := hy * 7
	return dy
}

// ToHumanYears function will take in dog years as a parameter and return the corresponding human years back
func ToHumanYears(dy int) int {
	hy := dy / 7
	return hy
}
