package main

import (
	"fmt"
)

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func AtoiBase(str string, base string) int {
	ans := 0
	for i := len(str) - 1; i >= 0; i-- {
		pow := (len(str) - 1) - i
		num := 0
		for j := 0; j <= len(base)-1; j++ {
			if base[j] == '+' || base[j] == '-' {
				return 0
			}
			if base[j] == str[i]{
				num = j
			}
		}
		piece := num * Power(len(base), pow)
		ans += piece
	}
	return ans
}

func Power(num int, pow int) int {
	if pow == 0 {
		return 1
	}
	if pow == 1 {
		return num
	}
	result := 1
	for i := 1; i <= pow; i++ {
		result *= num
	}
	return result
}
func Atoi(str string) int {
	res := 0
	for _, char := range str {
		if char >= '0' && char <= '9' {
			res = (res * 10) + int(char-'0')
		}
	}
	return res

}
