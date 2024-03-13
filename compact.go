package piscine

func Compact(ptr *[]string) int {
	count := 0
	for _, char := range *ptr {
		if char != "" {
			count++
		}
	}
	a := 0
	new := make([]string, count)
	for _, char := range *ptr {
		if char != "" {
			new[a] = char
			a++
		}
		*ptr = new
	}
	return count
}
