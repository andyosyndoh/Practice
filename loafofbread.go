package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output" + "\n"
	} else {
		new := ""
		x := 0
		for _, char := range str {
			if char%6 != 5 && char == ' ' {
				continue
			}
			if char%6 == 5 {
				new = new + ""
			} else {
				new = new + string(char)
			}
			x++
		}
		for i := len(new) - 1; i >= 0; i-- {
			if new[i] != ' ' {
				new = new[:i+1]
				break
			}
		}
		return new + "d\n"
	}
}
