package piscine

func BasicJoin(elems []string) string {
	q := ""
	for _, char := range elems {
		q = q + string(char)
	}
	return q
}
