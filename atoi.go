package piscine

func Atoi(str string) int {
	isneg := false
	if str[0] == '-' {
		str = str[1:]
		isneg = true
	}
	if str[0] == '+' {
		str = str[1:]
	}
	var num int

	for _, char := range str {
		if char >= '0' && char <= '9' {
			num = (num * 10) + int(char-'0')
		} else {
			return 0
		}
	}
	if isneg {
		num *= -1
	}
	return num
}
