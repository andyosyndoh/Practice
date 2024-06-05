package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 || len(os.Args) > 4 {
		return
	}
	str := os.Args[1]
	old := []rune(os.Args[2])
	new := []rune(os.Args[3])

	result := ""

	for _, ch := range str {
		if ch == old[0] {
			result += string(new[0])
		} else {
			result += string(ch)
		}
	}
	fmt.Println(result)
}
