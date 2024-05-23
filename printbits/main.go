package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	num := atoi(os.Args[1])

	for num > 0 {
		fmt.Print(num%2)
		num = num/2
	}
	fmt.Println()
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
