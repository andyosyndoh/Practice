package piscine

func AlphaCount(s string) int {
	str := []rune(s)
	count := 0
	for i := 0; i < len(s); i++ {
		if (str[i] >= 'A' && str[i] <= 'Z') || (str[i] >= 'a' && str[i] <= 'z') {
			count++
		}
	}
	return count
}
