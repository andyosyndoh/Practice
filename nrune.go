package piscine

func NRune(s string, n int) rune {
	j := []rune(s)
	if n < 1 || n > len(s) {
		return 0
	}
	return j[n-1]
}
