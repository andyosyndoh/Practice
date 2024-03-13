package piscine

type food struct {
	preptime int
  }
  
func FoodDeliveryTime(order string) int {
	menu := map[string]int {
		"burger": 15,
		"chips": 10,
		"nuggets": 12,
	}
	preptime, found := menu[order]
	if !found {
		return 0
	}
	return preptime

}
  