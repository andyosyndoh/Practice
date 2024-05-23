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
		return
	}
	if num == 3 {
		fmt.Println(false)
		return
	}
	i := 2 * 2
	for i <= num {
		if i == num {
			fmt.Println(true)
			break
		}
		i *= 2
		if i > num {
			fmt.Println(false)
			break
		}
		
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
