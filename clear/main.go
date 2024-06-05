package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("This is line 1")
    fmt.Println("This is line 2")
    fmt.Println("This is line 3")
    
    time.Sleep(2 * time.Second) // Wait for 2 seconds before clearing lines
    
    // Move the cursor up 2 lines
    fmt.Print("\033[3A")
    // Clear the current line
    fmt.Print("\033[3K")
    // Move to the beginning of the line
    fmt.Print("\033[G")
    // Clear the next line
    fmt.Print("\033[2K")
    // Move to the beginning of the line
    fmt.Print("\033[G")

    fmt.Println("This is the new line 1")
    fmt.Println("This is the new line 2")
    fmt.Println("This is the new line 3")
}
