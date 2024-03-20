package piscine

func Capitalize(s string) string {
	runes := []rune(s)

	for i := range runes {
		if runes[i] != runes[0] {
			if runes[i] == ' ' || runes[i] == '+' {
				if runes[i+1] >= 'a' && runes[i+1] <= 'z' {
					runes[i+1] = runes[i+1] - 'a' + 'A'
				} else if runes[i] >= 'A' && runes[i] <= 'Z' {
					runes[i] = runes[i] - 'A' + 'a'

				} else {
					runes[i+1] = runes[i+1]
				}
			}
		}
	}
	return string(runes)
}
