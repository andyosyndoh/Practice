package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	var new int
	for l := range args {
		new = l
	}

	for i := new; i > 0; i-- {
		for _, j := range args[i] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
