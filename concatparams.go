package piscine

func ConcatParams(args []string) string {
	var result []byte
	for i, arg := range args {
		result = append(result, arg...)
		if i < len(args)-1 {
			result = append(result, '\n')
		}
	}
	return string(result)
}
