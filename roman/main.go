package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	num, err := strconv.Atoi(os.Args[1])
	if err != nil || num <= 0 || num >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	ans, nums := romansnumber(num)
	fmt.Println(ans)
	fmt.Println(nums)
}

func romansnumber(num int) (string, string) {
	calculations := ""
	numerals := ""

	roman := []string{"M", "CM", "D", "XD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	calc := []string{"M", "(M-C)", "D", "(D-X)", "C", "(C-X)", "L", "(L-X)", "X", "(X-I)", "V", "(V-I)", "I"}

	for i, val := range numbers {
		for num >= val {
			num -= val
			numerals += roman[i]
			if len(calculations) > 0 {
				calculations += "+"
			}
			calculations += calc[i]
		}
	}

	return calculations, numerals
}
