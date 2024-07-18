package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <number>")
		return
	}

	input := os.Args[1]

	num, err := strconv.Atoi(input)
	if err != nil || num <= 0 || num >= 4000 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	roman := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	numbers := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}

	index := len(numbers) - 1
	var final string

	for num > 0 {
		for numbers[index] <= num {
			if num-numbers[index] == 0 {
				if roman[index] == "CM" || roman[index] == "CD" || roman[index] == "XC" || roman[index] == "XL" || roman[index] == "IX" || roman[index] == "IV" {
					fmt.Print("(", string(roman[index][1]), " - ", string(roman[index][0]), ")")
					final += roman[index]
					num -= numbers[index]
					break
				}
				final += roman[index]
				fmt.Print(roman[index])
				num -= numbers[index]
				break
			} else {
				if roman[index] == "CM" || roman[index] == "CD" || roman[index] == "XC" || roman[index] == "XL" || roman[index] == "IX" || roman[index] == "IV" {
					fmt.Print("(", string(roman[index][1]), " - ", string(roman[index][0]), ")", "+")
					final += roman[index]
					num -= numbers[index]
				} else {
					final += roman[index]
					fmt.Print(roman[index], "+")
					num -= numbers[index]
				}
			}
		}
		index--
	}

	fmt.Println()
	fmt.Println(final)
}
