package main

import (
	"fmt"
	"os"
	// "strconv"

	//"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 10 {
		return
	}
	num := Atoi(os.Args[1])

	if num <= 0 {
		return
	}

	for i := 1; i < 10; i++ {
		fmt.Printf("%v x %v = %v\n", i, num, i*num)
	}
}

func Atoi(str string) int {
	isneg := false
	if str[0] == '-' {
		str = str[1:]
		isneg = true
	}
	if str[0] == '+' {
		str = str[1:]
	}
	var num int

	for _, char := range str {
		if char >= '0' && char <= '9' {
			num = (num * 10) + int(char-'0')
		} else {
			return 0
		}
	}
	if isneg {
		num *= -1
	}
	return num
}