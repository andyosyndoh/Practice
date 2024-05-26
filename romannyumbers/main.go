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

	ans := ""
	num := 0
	
	for i, ch := range str {
		if str[i] == str[len(str)-1] {
			num = atoi(string(ch))
		} else {
			mod := 1
			num0 := len(str) - i - 1
			for i := 1; i <= num0; i++ {
				mod *=10
			}
			num = atoi(string(ch)) * mod
		}
		if st , ok := romanNumerals[num]; ok {
			ans += st 
		}
	}
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
