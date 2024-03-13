package piscine

func StringToIntSlice(str string) []int {
	var q []int
	for _, char := range str{
		p := int(char)
		q = append(q, p)
	}
	return q
}
