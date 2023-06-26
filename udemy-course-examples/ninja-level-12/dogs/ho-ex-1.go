// Package dog contains a subset of functions which allows users to operate on dog objects
package dogs

// Years function will take in human years as a parameter and return the corresponding dog years back
func Years(hy int) int {
	dy := hy * 7
	return dy
}
