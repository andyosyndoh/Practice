package main

import (
	"os"
)

func main() {
	if len(os.Args) == 1 || len(os.Args) > 4 {
		return
	}
	operators := []string{"*","/","-","+","%"}

	op := os.Args[2]

	cont := false

	for _, ch := range operators {
		if op == ch {
			cont = true
			break
		} else {
			cont = false
		}

	}
	if !cont {
		return
	}

	num1 := atoi(os.Args[1])
	num2 := atoi(os.Args[3])

	if num1 >= 9223372036854775807 || num1 <= -9223372036854775807 {
		return
	}
	if num2 >= 9223372036854775807 || num2 <= -9223372036854775807 {
		return
	}

	if op == "%" && num2 == 0 {
		os.Stdout.WriteString("No modulo by 0\n")
		return
	}
	if op == "/" && num2 == 0 {
		os.Stdout.WriteString("No division by 0\n")
		return
	}
	if op == "+" {
		result := itoa(num1+num2)
		os.Stdout.WriteString(result + "\n")
	}
	if op == "*" {
		result := itoa(num1*num2)
		os.Stdout.WriteString(result + "\n")
	}
	if op == "-" {
		result := itoa(num1-num2)
		os.Stdout.WriteString(result + "\n")
	}
	if op == "/" {
		result := itoa(num1/num2)
		os.Stdout.WriteString(result + "\n")
	}
	if op == "%" {
		result := itoa(num1%num2)
		os.Stdout.WriteString(result + "\n")
	}
	
	
}

func atoi(m string) int {
	// d := []rune(m)
	for _, ch := range m {
		if ch >= '0' && ch <= '9' {
			continue
		} else {
			os.Exit(0)
		}
	}
	if m == "0" {
		return 0
	}
	n := 0
	for _, ch := range m {
		n = n*10 + int(ch-'0')
	}
	return n 
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	isneg := false
	if n < 0 {
		isneg = true
		n = -n
	}
	res := ""
	for n > 0 {
		digit := n%10
		res = string('0'+digit) + res 
		n /= 10
	}
	if isneg {
		res = "-" + res
	}
	return res
}
