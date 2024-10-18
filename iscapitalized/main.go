package main

import "fmt"

func main() {
	fmt.Println(IsCapitalized("Hello and There"))
	fmt.Println(IsCapitalized("Hello And There"))
}

func IsCapitalized(str string) bool {
	if str == "" || Small(rune(str[0])) {
		return false
	}
	for i := 1; i < len(str); i++ {
		if str[i] == ' ' && Small(rune(str[i+1])) {
			return false
		}
	}
	return true
}

func Small(r rune) bool {
	return r >= 'a' && r <= 'z'
}
