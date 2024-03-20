package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	args := os.Args
	new := args[0]

	name := []rune(new)
	for  ch := range name {
		z01.PrintRune(name[ch])
	}
	z01.PrintRune('\n')
}
