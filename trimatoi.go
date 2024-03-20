package piscine

func TrimAtoi(s string) int {
	var neg bool
	var newer int

	for _, ch := range s {
		if ch == '-' && newer == 0 {
			neg = true
		}
		if ch >= '1' && ch <= '9' {
			newer = newer*10 + int(ch-'0')
		}
	}
	if neg {
		return -newer
	}
	return newer
}
