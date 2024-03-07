package piscine

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if (s[i] < 'a' || s[i] > 'z') && (s[i] < 'A' || s[i] > 'Z') && (s[i] < '1' || s[i] > '0') {
			return false
		}
	}
	return true
}
