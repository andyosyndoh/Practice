package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {

		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	for _, ch := range Itoa(len(arg)) {
		z01.PrintRune(ch)
	}
	z01.PrintRune('\n')

}

func Itoa(s int) string {
	if s == 0 {
		return "0"
	}

	num := ""
	for s > 0 {
		digit := s % 10
		num = string(digit+'0') + num
		s /= 10
	}
	return num
}
