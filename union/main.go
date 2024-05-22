package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 3 || len(os.Args) == 1 {
		return
	}
	args := ""

	for _, str := range os.Args[1:] {
		args += str
	}

	seen := make(map[rune]bool)
	result := ""

	for _, ch := range args {
		if !seen[ch] {
			seen[ch] = true
			result += string(ch)
		}
	}

	fmt.Println(result)
}
