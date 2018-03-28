package plot

import "strconv"

const maxDigits = 7

// Ff formats a float
func Ff(x float64) []rune {
	minExact := strconv.FormatFloat(x, 'g', -1, 64)
	fixed := strconv.FormatFloat(x, 'g', maxDigits, 64)
	if len(minExact) < len(fixed) {
		return []rune(minExact)
	}
	return []rune(fixed)
}
