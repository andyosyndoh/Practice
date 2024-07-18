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
	res := 0
	result := ""

	bits := "01"

	for num > 0 {
		res = (num % len(bits))
		result = string(bits[res]) + result
		num = num / len(bits)
	}

	zeros := ""
	for i := 0; i < 8-len(result); i++ {
		zeros += "0"
	}
	fmt.Println(zeros + result)
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
