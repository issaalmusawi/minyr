package main

import (
	"fmt"
	"os"
)

func main() {
    // Open the file for reading
    file, err := os.Open("kjevik-temp-celsius-20220318-20230318.csv")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()

    // Read the file
    buffer := make([]byte, 85)
    bytesRead, err := file.Read(buffer)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Print the file content
    fmt.Println(string(buffer[:bytesRead]))
}

