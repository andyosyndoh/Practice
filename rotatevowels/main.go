package main

import (
	"fmt"
	"os"
	//
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println()
		return
	}
	args := ""

	for i, str := range os.Args[1:] {
		if i != 0 {
			args += " " + str
		} else {
			args += str
		}
	}
	list := []rune{}
	for _, ch := range args {
		if Isvowel(ch) {
			list = append(list, (ch))
		}
	}
	result := ""

	count := 1
	for _, ch := range args {
		if Isvowel(ch) {
			ch = list[len(list)-count]
			result += string(ch) 
			count++
		} else {
			result += string(ch)
		}
		
	}
	fmt.Println(result)

}

func Isvowel(s rune) bool {
	vowels := "aeiouAEIOU"
	for _, ch := range vowels {
		if s == ch {
			return true
		}
	}
	return false
}
