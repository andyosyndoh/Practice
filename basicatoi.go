package piscine

func BasicAtoi(s string) int {
	new := 0
	for _, val := range s {
		new = new*10 + int(val-'0')
	}
	return new
}
