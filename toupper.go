package piscine

func ToUpper(s string) string {
	capitalized := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' || s[i] <= 'z' {
			capitalized += string(s[i] - 'a')
		} 
	}
	return capitalized
}
