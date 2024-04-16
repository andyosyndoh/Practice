package piscine

func BasicAtoi2(s string) int {
	for _, ch := range s {
		if ch >= '1' && ch <= '9' && ch != ' ' {
			new := 0
			for _, val := range s {
				new = new*10 + int(val-'0')
			}
			return new
		} else {
			return 0
		}
	}
	return 0
}
