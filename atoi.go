package piscine

// func Atoi(str string) int {
// 	isneg := false
// 	if str[0] == '-' {
// 		str = str[1:]
// 		isneg = true
// 	}
// 	if str[0] == '+' {
// 		str = str[1:]
// 	}
// 	var num int

// 	for _, char := range str {
// 		if char >= '0' && char <= '9' {
// 			num = (num * 10) + int(char-'0')
// 		} else {
// 			return 0
// 		}
// 	}
// 	if isneg {
// 		num *= -1
// 	}
// 	return num
// }

func Atoi(s string) int {
	isneg := false
	if s[0] == '-' {
		isneg = true
		s = s[1:]
	}
	if s[0] == '+' {
		isneg = false
		s = s[1:]
	}
	var num int
	for _, val := range s {
		if val >= '0' && val <= '9' {
			num = (num * 10) + int(val-'0')
		} else {
			return 0
		}
	}
	if isneg {
		num *= -1
	}
	return num
}
