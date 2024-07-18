package piscine

func StrRev(s string) string {
	x := []rune(s)
	var new []rune
	for i := len(s) - 1; i >= 0; i-- {
		new = append(new, x[i])
	}
	return string(new)
}
