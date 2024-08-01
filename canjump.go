package piscine

func CanJump(arr []uint) bool {
	sum := 0
	for _, ch := range arr {
		sum += int(ch)
	}
	if sum == 0 {
		return true
	}
	i := 0
	for i < len(arr) {
		if i == len(arr)-1 {
			return true
		}
		if arr[i] == uint(0) {
			return false
		}
		i = int(arr[i]) + i
	}
	return false
}
