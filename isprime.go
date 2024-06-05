package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

// package piscine

// func IsPrime(num int) bool {
// 	if num <= 1 {
// 		return false
// 	}
// 	if num <= 3 {
// 		return true
// 	}
// 	if num%2 == 0 || num%3 == 0 {
// 		return false
// 	}
// 	for i:= 5; i*i <= num ; i+=6 {
// 		if num%i == 0 || num %(i+2) == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }
