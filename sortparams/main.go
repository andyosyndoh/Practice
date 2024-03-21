package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	for i := 0; i < len(arguments)-1; i++ {
		for k := i + 1; k < len(arguments); k++ {
			if arguments[i] > arguments[k] {
				arguments[i], arguments[k] = arguments[k], arguments[i]
			}
		}

	}
	for j := 0; j < len(arguments); j++ {
		for _, ch := range arguments[j] {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
	
}
