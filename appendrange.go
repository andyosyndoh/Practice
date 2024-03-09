package piscine

func AppendRange(min, max int) []int {
	var answer []int
	for i := min-1; i < max-1; i++ {
		answer = append(answer, i+1)
	}
	return answer
}
