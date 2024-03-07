package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		z01.PrintRune(0)
	}
	digits := make([]int, 10)
	result := make([]rune, 0, 10)
	for n > 0 {
		digits[n%10]++
		n /= 10
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < digits[i]; j++ {
			result = append(result, rune(i)+'0')
		}
	}
	for _, digit := range result {
		z01.PrintRune(digit)
	}
}
