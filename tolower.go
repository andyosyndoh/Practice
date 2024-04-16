package piscine

func ToLower(s string) string {
	new := make([]rune, len(s))

	for i, char := range s {
		if char >= 'A' && char <= 'Z' {
			new[i] = char - 'A' + 'a'
		} else {
			new[i] = char
		}
	}
	return string(new)
}
