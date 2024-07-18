package piscine

func SplitWhiteSpaces(s string) []string {
	var new []string
	old := ""
	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if old != "" {
				new = append(new, old)
				old = ""
			}
		} else {
			old += string(char)
		}
	}
	if old != "" {
		new = append(new, old)
	}
	return new
}
