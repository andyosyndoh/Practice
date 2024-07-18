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

	if num > 0 && num&(num-1) == 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
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
