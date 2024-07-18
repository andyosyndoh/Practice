package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func PrintStr(s string) {
	runes := []rune(s)
	for _, char := range runes {
		z01.PrintRune(char)
	}
}

func setPoint(ptr *point) {
	ptr.x = 43
	ptr.y = 22
}

func Output(nb int) {
	numdiv := '0'
	nummod := '0'
	for i := 1; i <= nb/10; i++ {
		numdiv++
	}
	z01.PrintRune(numdiv)
	for j := 2; j <= nb%10; j++ {
		nummod++
	}
	z01.PrintRune(nummod)
}

func main() {
	points := &point{}

	setPoint(points)
	x := points.x
	y := points.y

	PrintStr("x = ")
	Output(x)
	PrintStr(", y = ")
	Output(y)
	z01.PrintRune('\n')
}
