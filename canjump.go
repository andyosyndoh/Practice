package piscine

func CanJump(arr []uint) bool {
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
