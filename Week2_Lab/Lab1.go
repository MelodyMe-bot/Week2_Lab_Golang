//go:build ignore

package main

import "fmt"

func main() {
	var x int = 15
	var y int = 35

	x = y
	fmt.Println(x)

	// add
	x += y
	fmt.Println(x)

	// substract
	x -= y
	fmt.Println(x)

	//multiply
	x *= y
	fmt.Println(x)

	//Division
	x /= y
	fmt.Println(x)

	//Modulus
	x %= y
	fmt.Println(x)

}
