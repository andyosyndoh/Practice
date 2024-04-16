package piscine

func Map(f func(int) bool, a []int) []bool {
	length := 1
	for l := range a {
		length = l + 1
	}
	b := make([]bool, length)
	for i := range a {
		b[i] = f(a[i])
	}
	return b
}
