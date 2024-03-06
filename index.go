package piscine

func Index(s string, toFind string) int {
	a := []rune(s)
	b := []rune(toFind)

	if len(s) == 0 || len(toFind) == 0 {
		return 0
	}
	for i := 0; i <= len(s); i++ {
		if a[i] == b[0] {
			return i
		} else {
			return -1
		}
	}
	return 0
}
