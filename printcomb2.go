package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	next := false
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '8'; j++ {
			for k := '0'; k <= '9'; k++ {
				for l := '1'; l <= '9'; l++ {
					if j < l && j <= k {
						if next {
							z01.PrintRune(',')
							z01.PrintRune(' ')

						}
						next = true
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(l)
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
