package piscine

func MakeRange(min, max int) []int {
	new := []int{}
	for i := min; i < max; i++ {
		new = append(new, i)
	}
	return new
}
