package main

import "fmt"

func ReverseBits(octet byte) byte {
	var ans byte = 0
	for i := 1; i <= 8; i++ {
		bit := (octet >> i) & 1
		ans = (ans << 1) | bit
	}
	return ans
}

func main() {
	// Example usage
	inputByte := byte(8) // Example byte: 01000001
	swappedByte := ReverseBits(inputByte)
	fmt.Printf("Original byte: %08b\n", inputByte)
	fmt.Printf("Swapped byte:  %08b\n", swappedByte)
}
