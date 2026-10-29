//go:build ignore

package main

import "fmt"

func main() {
	x := 25
	y := 30

	if x > 0 && y > 0 {
		fmt.Println("Both number are positive")
	} else {
		fmt.Println("One number are negative!")
	}

	if x > y || y > x {
		fmt.Println("One number are greater!")
	}

	if x != y {
		fmt.Println("Both number are not equal each other!")
	}

}
