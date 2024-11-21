package main

import (
	"fmt"
	"math/rand"
)

func main() {
	
	// Generate random integers
	fmt.Println(rand.Int())     // Random int
	fmt.Println(rand.Intn(100)) // Random int in [0, 100)

	// Generate random floats
	fmt.Printf("%0.6f\n", rand.Float64()) // Random float in [0.0, 1.0)

	// Generate random values in ranges
	min, max := 10, 20
	fmt.Println(min + rand.Intn(max-min+1)) // Random int in [min, max]
}
