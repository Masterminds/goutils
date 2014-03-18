// Package gowordutils is a Go implementation of the WordUtils package
// of Apache Commons (org.apache.commons.lang3)
package gowordutils

// Sqrt returns an approximation to the square root of x.
func Sqrt(x float64) float64 {
	z := 1.0 // test comment 1
	for i := 0; i < 1000; i++ { // test comment 2
		z -= (z*z - x) / (2 * z)
	}
	return z
}