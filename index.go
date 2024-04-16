package piscine

func Index(s string, toFind string) int {
	h := len(s)
	k := len(toFind)
	for i := 0; i <= h-k; i++ {
		if s[i:i+k] == toFind {
			return i
		}
	}
	return -1
}
