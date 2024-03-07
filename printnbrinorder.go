package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}
	digits := make([]int, 10)
	result := make([]rune, 0, 10)
	for n > 0 {
		digit := n%10
		digits[digit]++
		n /= 10
	}
	
	for i :=0; i < 10; i++ {
		if i == 0 {
			result = append(result, '0')
		}
		for j := 0; j < digits[i]; j++ {
			result = append(result, rune(i)+'0')
		}
	}
	for _, digit := range result {
		z01.PrintRune(digit)
	}
	z01.PrintRune(0)
}
