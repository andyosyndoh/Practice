package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, ch := range tab {
		if f(ch) {
			count++
		}
	}
	return count
}
