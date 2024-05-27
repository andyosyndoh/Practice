package main

import (
	"fmt"
	"os"
)

var romanNumerals = map[int]string{
	1:    "I",
	2:    "II",
	3:    "III",
	4:    "IV",
	6:    "VI",
	5:    "V",
	7:    "VII",
	8:    "VIII",
	9:    "IX",
	10:   "X",
	20:   "XX",
	30:   "XXX",
	40:   "XL",
	50:   "L",
	60:   "LX",
	70:   "LXX",
	80:   "LXXX",
	90:   "XC",
	100:  "C",
	200:  "CC",
	300:  "CCC",
	400:  "CD",
	500:  "D",
	600:  "DC",
	700:  "DCC",
	800:  "DCCC",
	900:  "CM",
	1000: "M",
	2000: "MM",
	3000: "MMM",
}

func main() {
	str := os.Args[1]

	answer := []string{}
	ans := ""
	num := 0
	final := ""

	for i, ch := range str {
		if i == len(str)-1 {
			num = atoi(string(ch))
		} else {
			state := 1
			num0 := len(str) - i - 1
			for i := 1; i <= num0; i++ {
				state *= 10
			}
			num = atoi(string(ch)) * state
		}
		if st, ok := romanNumerals[num]; ok {
			ans += st
		}
		sli := number(num)
		for i := 0; i < len(sli); i++ {
			answer = append(answer, sli[i])
		}

	}

	for i := 0; i < len(answer); i++ {
		if i == 0 {
			final += answer[i]
		} else {
			final += "+" + answer[i]
		}
	}

	fmt.Println(final)
	fmt.Println(ans)
}

func atoi(m string) int {
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

func number(num int) []string {
	ans := []string{}
	if num >= 1 && num <= 3 {
		for i := 1; i <= num; i++ {
			ans = append(ans, "I")
		}
	} else if num == 4 {
		ans = append(ans, "(V-I)")
	} else if num == 5 {
		ans = append(ans, "V")
	} else if num >= 6 && num <= 8 {
		ans = append(ans, "V")
		for i := 6; i <= num; i++ {
			ans = append(ans, "I")
		}
	} else if num == 9 {
		ans = append(ans, "(X-I)")
	} else if num >= 10 && num <= 30 {
		for i := 10; i <= num; i += 10 {
			ans = append(ans, "X")
		}
	} else if num == 40 {
		ans = append(ans, "(L-X)")
	} else if num == 50 {
		ans = append(ans, "L")
	} else if num >= 60 && num <= 80 {
		ans = append(ans, "L")
		for i := 60; i <= num; i += 10 {
			ans = append(ans, "X")
		}
	} else if num == 90 {
		ans = append(ans, "(C-X)")
	} else if num >= 100 && num <= 300 {
		for i := 100; i <= num; i = i + 100 {
			ans = append(ans, "C")
		}
	} else if num == 400 {
		ans = append(ans, "(D-C)")
	} else if num == 500 {
		ans = append(ans, "D")
	} else if num >= 600 && num <= 800 {
		ans = append(ans, "D")
		for i := 600; i <= num; i++ {
			ans = append(ans, "C")
		}
	} else if num == 900 {
		ans = append(ans, "(M-C)")
	} else if num >= 1000 && num <= 3000 {
		for i := 1000; i <= num; i = i + 1000 {
			ans = append(ans, "M")
		}
	}
	return ans
}
