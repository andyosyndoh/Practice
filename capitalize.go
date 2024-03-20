package piscine

func Capitalize(s string) string {
	runes := []rune(s)

	for i := range runes {
		if runes[i] == ' ' ||  runes[i] == '+' {
            if runes[i+1] >= 'a' && runes[i+1] <= 'z' {
              runes[i+1] = runes[i+1] - 'a' + 'A'
            } else {
              runes[i+1] = runes[i+1]
            }
          
        }
    }
	return string(runes)
}


