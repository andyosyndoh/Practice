package piscine

func atoi(s string) int {
	isneg := false
	if s[0] == '-' {
		isneg = true
	}
	if s[0] == '+' {
		isneg = false
	}
	var num int
	for _, val := range s {
		if val >= '0' && val <= '9' {
			num = (num * 10) + int(val-'0')
		}
	}
	if !isneg {
		num *= -1
	}
	return num
}

