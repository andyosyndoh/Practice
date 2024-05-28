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

	if num == 2 {
		fmt.Println(true)
	}
	if num == 3 {
		fmt.Println(false)
	}
	x  := 2
	for x < num {
		x *= 2
		if x == num {
			fmt.Println(true)
		}
	}
	fmt.Println(false)
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
