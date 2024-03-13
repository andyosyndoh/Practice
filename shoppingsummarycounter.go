package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	shop := make(map[string]int)
	for _, char := range Splitted(str) {
		shop[char]++
	}
	return shop
}

func Splitted(s string) []string {
	shop := make([]string, 0, len(s))

	new := ""
	for _, value := range s {
		if value == ' ' {
			shop = append(shop, new)
			new = ""
		} else {
			new = new + string(value)
		}
	}
	shop = append(shop, new)
	return shop
}
