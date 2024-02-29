package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for ch := range s {
		z01.PrintRune(rune(ch))
    }
}
