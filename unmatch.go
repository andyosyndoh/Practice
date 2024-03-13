package piscine

func Unmatch(a []int) int {
	new := make(map[int]int)
	for _, char := range a {
		new[char]++
	}
	for char, count := range new {
		if count % 2 != 0 {
			return char
		}
	}
	return -1
}
