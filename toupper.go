package piscine

func ToUpper(s string) string {
	new := ""
	for i := 1; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			new = string(32) + string(s)
		}
	}
	return new
}
