package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var s []int

	if n == 0 {
		s = append(s, 0)
	}
	isnegative := false
	if n < 0 {
		isnegative = true
		n = -n
	}
	for n > 0 {
		digit := n%10
		s = append(s, digit)
		n = n/10
	}
	if isnegative {
		z01.PrintRune('-')
	}

	for i := len(s)-1; i >= 0 ; i-- {
		z01.PrintRune(rune('0' + s[i]))
	}
}
