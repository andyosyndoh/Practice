package main

import "fmt"

func main() {
	text := "Hello, world!"
	width := 20                          // Total width of alignment
	fmt.Printf("%-[1]*s\n", width, text) // Left align
	fmt.Printf("%[1]*s\n", width, text)  // Right align
	fmt.Printf("%[1]*s\n", -width, text) // Center align
}
