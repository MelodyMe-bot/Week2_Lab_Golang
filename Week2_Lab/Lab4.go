//go:build ignore

package main

import "fmt"

func add(a, b int) int { return a + b }
func sub(a, b int) int { return a - b }
func mul(a, b int) int { return a * b }
func div(a, b int) int { return a / b }
func mod(a, b int) int { return a % b }

func main() {
	for {
		fmt.Println("\n===== Mini Calculator =====")
		fmt.Println("1) Add  2) Sub  3) Mul  4) Div  5) Mod  6) Exit")
		fmt.Print("Choose: ")

		var choice string
		fmt.Scan(&choice)

		if choice == "6" {
			fmt.Println("Goodbye!")
			break
		}

		if choice != "1" && choice != "2" && choice != "3" && choice != "4" && choice != "5" {
			fmt.Println("Error: invalid choice")
			continue
		}

		var a, b int
		fmt.Print("Enter a: ")
		fmt.Scan(&a)
		fmt.Print("Enter b: ")
		fmt.Scan(&b)

		if choice == "1" {
			fmt.Println("Result:", add(a, b))
		} else if choice == "2" {
			fmt.Println("Result:", sub(a, b))
		} else if choice == "3" {
			fmt.Println("Result:", mul(a, b))
		} else if choice == "4" {
			if b == 0 {
				fmt.Println("Error: division by zero")
			} else {
				fmt.Println("Result:", div(a, b))
			}
		} else if choice == "5" {
			if b == 0 {
				fmt.Println("Error: division by zero")
			} else {
				fmt.Println("Result:", mod(a, b))
			}
		}
	}
}
