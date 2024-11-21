package main

import "fmt"

func main() {
	fmt.Println(ItoaBase(10, 2))
	fmt.Println(ItoaBase(255, 16))
	fmt.Println(ItoaBase(-42, 4))
	fmt.Println(ItoaBase(123, 10))
	fmt.Println(ItoaBase(0, 8))
	fmt.Println(ItoaBase(255, 2))
	fmt.Println(ItoaBase(-255, 16))
	fmt.Println(ItoaBase(15, 16))
	fmt.Println(ItoaBase(10, 4))
	fmt.Println(ItoaBase(255, 10))
}

func ItoaBase(n int, b int) string {
	if n == 0 {
        return "0"
    }
	str := "0123456789ABCDEF"

	var result string
	var sign string

	if n < 0 {
		sign = "-"
		n = -n
	}

	num := uint(n)
	base := uint(b)

	for num > 0 {
		remainder := num % base
		result = string(str[remainder]) + result
		num /= base
	}

	return sign + result
}
