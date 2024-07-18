package piscine

func BasicJoin(elems []string) string {
	q := ""
	for _, ch := range elems {
		q = q + string(ch)
	}
	return q
}
