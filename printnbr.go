package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var s []int

	if n == 0 {
		s = append(s, 0)
	}
	for i := 1; n > 0; i++ {
		s = append(s, n%10)
		n /= 10
	}
	for i := range s {
		z01.PrintRune(rune('0' + s[i]))
	}
}
