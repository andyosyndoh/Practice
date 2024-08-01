package main

import (
	"fmt"
)

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}

func CamelToSnakeCase(s string) string {
	str := ""
	for i, ch := range s {
		if IsCap(ch) && IsCap(rune(s[i+1])) {
			return s
		}
		if IsCap(ch) && i != 0{
			str += "_" + string(ch)
		} else {
			str += string(ch)
		}

	}
	return str
}

func IsCap(d rune) bool {
	alph := "QWERTYUIOPLKJHGFDSAZXCVBNM"
	for _ , let := range alph {
		if d == let {
			return true
		}
	}
	return false
}
