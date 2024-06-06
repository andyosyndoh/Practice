package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		return
	}
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
		num := atoi(str)
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

func atoi(s string) int {
	if s == "0" {
		return 0
	}
	num := 0
	for _, ch := range s {
		num = num*10 + int(ch-'0')
	}
	return num
}
