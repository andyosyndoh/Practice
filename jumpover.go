package piscine

func JumpOver(str string) string {
	new := ""
	if len(str) > 3 {
		for a, char := range str {
			if a%3 == 2 {
				new = new + string(char)
			}
		}
	}
	return new + "\n"
}
