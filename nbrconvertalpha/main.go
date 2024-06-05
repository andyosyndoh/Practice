package main

import (
	"os"

	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	args := []string{}
	up := false
	if os.Args[1] == "--upper" {
		up = true
		args = os.Args[2:]
	} else {
		args = os.Args[1:]
	}

	argnum := []int{}
	for _, str := range args {
		num := piscine.Atoi(str)
		if num < 1 || num > 26 {
			num = 0
		}
		argnum = append(argnum, num)
	}

	alph := "abcdefghijklmnopqrstuvwxyz"
	new := ""

	for _, num := range argnum {
		if num == 0 {
			new += " "
		} else {
			new += string(alph[num-1])
		}
	}

	for _, ch := range new {
		if ch == ' ' {
			z01.PrintRune(ch)
		} else {
			if up {
				z01.PrintRune(ch - 32)
			} else {
				z01.PrintRune(ch)
			}
		}
	}
	z01.PrintRune('\n')
}
