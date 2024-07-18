package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	num1 := atoi(os.Args[1])
	num2 := atoi(os.Args[2])
	ref := 0
	if num1 > num2 {
		ref = num2
	} else {
		ref = num1
	}

	for i := ref - 1; i >= 1; i-- {
		if num1%i == 0 && num2%i == 0 {
			fmt.Println(i)
			break
		}
	}
}

func atoi(m string) int {
	// d := []rune(m)
	for _, ch := range m {
		if ch >= '0' && ch <= '9' {
			continue
		} else {
			os.Exit(0)
		}
	}
	if m == "0" {
		return 0
	}
	n := 0
	for _, ch := range m {
		n = n*10 + int(ch-'0')
	}
	return n
}
