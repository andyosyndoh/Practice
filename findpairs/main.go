package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}
	nums := os.Args[1]
	target := os.Args[2]

	if len(nums) < 3 || nums[0] != '[' || nums[len(nums)-1] != ']' {
		fmt.Println("Invalid input.")
		return
	} else {
		nums = nums[1 : len(nums)-1]
	}

	numsslice := Split(nums, ", ")
	var numbers []int
	for _, num := range numsslice {
		conv, err := strconv.Atoi(num)
		if err != nil {
			fmt.Printf("Invalid number: %v\n", num)
			return
		}
		numbers = append(numbers, conv)
	}
	conv, err := strconv.Atoi(target)
	if err != nil {
		fmt.Printf("Invalid target sum.\n")
		return
	}

	var answers [][]int
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			var ans []int
			if numbers[i]+numbers[j] == conv {
				ans = append(ans, i)
				ans = append(ans, j)
			}
			if len(ans) == 2 {
				answers = append(answers, ans)
			}
		}
	}
	if len(answers) == 0 {
		fmt.Println("No pairs found.")
		return
	}
	fmt.Printf("Pairs with sum %v: %v\n", conv, answers)
}

func Split(str, sep string) []string {
	var result []string
	start := 0
	for i := 0; i < len(str); i++ {
		if i+len(sep) < len(str) {
			if str[i:i+len(sep)] == sep {
				result = append(result, str[start:i])
				start = i + len(sep)
				i = i + len(sep)
			}
			if i == len(str)-1 {
				result = append(result, str[start:])
			}
		}
	}

	return result
}
