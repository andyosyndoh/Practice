package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	new := args[0][2:]

	name := []rune(new)
	for ch := range name {
		z01.PrintRune(name[ch])
	}
	z01.PrintRune('\n')
}
