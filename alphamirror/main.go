package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1]
	for _, ch := range arg {
		if ch >= 'a' && ch <= 'z' {
			ch = 'z' - (ch - 'a')
			z01.PrintRune(ch)
		}
		if ch >= 'A' && ch <= 'Z' {
			ch = 'Z' - (ch - 'A')
			z01.PrintRune(ch)
		}
	}
	z01.PrintRune('\n')
}
