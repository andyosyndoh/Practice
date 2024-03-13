package piscine

func ReverseMenuIndex(menu []string) []string {
	new := make([]string, len(menu))
	for i := 0; i < len(menu); i++ {
		new[len(menu)-1-i] = new[len(menu)-1-i] + menu[i]
	}
	return new
}
