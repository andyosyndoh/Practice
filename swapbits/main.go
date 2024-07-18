package main

import "fmt"

func SwapBits(octet byte) byte {
	// Perform the bit swap
	return (octet<<4 | octet>>4)
}

func main() {
	// Example usage
	inputByte := byte(0b01000001) // Example byte: 01000001
	swappedByte := SwapBits(inputByte)
	fmt.Printf("Original byte: %08b\n", inputByte)
	fmt.Printf("Swapped byte:  %08b\n", swappedByte)
}
