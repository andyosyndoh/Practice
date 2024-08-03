package main

import (
	"fmt"
	"os"
	"strings"
)

const optionsString = "abcdefghijklmnopqrstuvwxyz"

func main() {
	args := os.Args[1:]

	if len(args) == 0 || (len(args) > 0 && strings.Contains(args[0], "h")) {
		fmt.Println("options: " + optionsString)
		return
	}

	var options int32 = 0

	for _, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			fmt.Println("Invalid Option")
			return
		}

		for _, char := range arg[1:] {
			if strings.ContainsRune(optionsString, char) {
				options |= (1 << (char - 'a'))
			} else {
				fmt.Println("Invalid Option")
				return
			}
		}
	}

	fmt.Printf("%08b %08b %08b %08b\n",
		(options>>24)&0xFF,
		(options>>16)&0xFF,
		(options>>8)&0xFF,
		options&0xFF)
}
