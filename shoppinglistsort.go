package piscine

func ShoppingListSort(slice []string) []string {
	for j := 0; j < len(slice); j++ {
		for i := j + 1; i < len(slice); i++ {
			if len(slice[j]) > len(slice[i]) {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}
