package piscine

func Join(strs []string, sep string) string {
	q := ""
	for i, char := range strs {
		q = q + string(char)
		if i != len(strs)-1 {
			q = q + sep
		}
	}
	return q
}
