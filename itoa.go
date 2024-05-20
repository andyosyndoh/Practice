package piscine

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	isnegative := false
	if n < 0 {
		isnegative = true
		n = n * -1
	}
	new := ""
	for n > 0 {
		digit := n%10
		new = string('0' + digit) + new
		n = n/10
	}
	if isnegative {
		new = "-" + new
	}
	return new
}