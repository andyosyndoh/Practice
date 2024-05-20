package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 3 || len(os.Args) == 1 {
		return
	}

	arg1 := os.Args[1]
	arg2 := os.Args[2]

	result := wdmatch(arg1,arg2)

	fmt.Print(result)

}

func wdmatch(s, w string) string {
	var i,j int
	for i < len(s) && j < len(w) {
		if s[i] == w[j] {
			i++
		}
		j++
	}
	if i == len(s) {
		return s + "\n"
	}
	return ""
}