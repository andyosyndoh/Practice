package piscine

func ToUpper(s string) string {
	new := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			new = string(s[i]- 32) + string(s)
		} 
	}
	return new
}
