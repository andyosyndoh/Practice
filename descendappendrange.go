package piscine

func DescendAppendRange(max, min int) []int {
	var answer = []int{}
	if max < min {
		return answer
	}
	for i := max; i > min; i-- {
		answer = append(answer, i)
	}
	return answer
}
