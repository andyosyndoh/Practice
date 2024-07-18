package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	str := os.Args[1]

	vows := []byte{'a', 'e', 'i', 'o', 'u'}
	for i := 0; i < len(str); i++ {
		for _, vow := range vows {
			if str[0] == vow {
				str = str + "ay"
				fmt.Println(str)
				return
			} else {
				if str[i] == vow {
					str = str[i:] + str[:i] + "ay"
					fmt.Println(str)
					return
				}
			}
		}
	}
	fmt.Println("No vowels")
}
