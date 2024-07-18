package piscine

func ToUpper(s string) string {
	new := make([]rune, len(s))

	for i, char := range s {
		if char >= 'a' && char <= 'z' {
			new[i] = char - 'a' + 'A'
		} else {
			new[i] = char
		}
	}
	return string(new)
}
