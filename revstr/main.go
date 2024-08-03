package main

import (
	"fmt"
	"os"
)

func main() {
	str := os.Args[1]

	newstr := []string{}
	result := ""

	start := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			newstr = append(newstr, str[start:i])
			start = i + 1
		}
		if i == len(str)-1 {
			newstr = append(newstr, str[start:])
		}
	}

	if len(newstr) > 1 {
		for i := len(newstr) - 1; i > 0; i-- {
			result += newstr[i] + " "
		}
	} else {
		result = newstr[0] + " "
	}

	fmt.Println(result[:len(result)-1])
}
