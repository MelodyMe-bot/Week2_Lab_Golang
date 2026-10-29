package main

import (
	"bufio"
	"fmt"
	"os"
)

func xorEncrypt(text string, key byte) string {
	result := ""
	for i := 0; i < len(text); i++ {
		result += string(text[i] ^ key)
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter message: ")
	scanner.Scan()
	text := scanner.Text()

	fmt.Print("Enter key (1 character): ")
	scanner.Scan()
	keyText := scanner.Text()

	if text == "" || keyText == "" {
		fmt.Println("Error: message and key cannot be empty")
		return
	}
	key := keyText[0]

	cipher := xorEncrypt(text, key)

	fmt.Print("Ciphertext (hex): ")
	for i := 0; i < len(cipher); i++ {
		fmt.Printf("%02x ", cipher[i])
	}
	fmt.Println()

	// Same function decrypts
	plain := xorEncrypt(cipher, key)
	fmt.Println("Decrypted:", plain)
}
