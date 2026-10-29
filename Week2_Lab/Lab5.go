//go:build ignore

package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter text: ")
	scanner.Scan()
	text := scanner.Text()

	// Binary
	fmt.Print("Binary: ")
	for i := 0; i < len(text); i++ {
		fmt.Printf("%08b ", text[i])
	}
	fmt.Println()

	// Hex
	fmt.Print("Hex:    ")
	for i := 0; i < len(text); i++ {
		fmt.Printf("%02x ", text[i])
	}
	fmt.Println()

	// Base64
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	fmt.Println("Base64:", encoded)
}
