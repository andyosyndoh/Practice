package piscine

func IsAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		if (s[i] < 'a' || s[i] > 'z') && (s[i] < 'A' || s[i] > 'Z') {
			return false
		}
	}
	return true
}
