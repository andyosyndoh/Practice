package piscine

func Unmatch(a []int) int {
	for _, char := range a {
		i := 0
		for _, b := range a {
			if b == char {
				i++
			}
		}
		if i == 1 || i%2 == 1 {
			return char
		}
	}
	return -1
}
