package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println()
		return
	}

	str := os.Args[1]
	for str[0] == ' ' {
		str = str[1:]
	}
	for str[len(str)-1] == ' ' {
		str = str[:len(str)-1]
	}
	for i := 0; i < len(str); i++ {
		for str[i] == ' ' && str[i+1] == ' ' {
			str = str[:i+1] + str[i+2:]
		}
	}

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			str = str[i+1:] + " " + str[:i]
			break
		}
	}
	fmt.Println(str)
}
