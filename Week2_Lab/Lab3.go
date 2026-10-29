//go:build ignore

package main

import "fmt"

func myXOR(a, b int) {
	fmt.Printf("XOR(A,B) = %d ^ %d = %d\n", a, b, a^b)
}

func myOR(a, b int) {
	fmt.Printf("OR(A,B) = %d | %d = %d\n", a, b, a|b)
}

func myAND(a, b int) {
	fmt.Printf("AND(A,B) = %d & %d = %d\n", a, b, a&b)
}

func myNOT(a, b int) {
	fmt.Printf("%d = %d\n", a, ^a)
	fmt.Printf("%d = %d\n", b, ^b)

}

func myShift(a, b int) {
	fmt.Printf("Left Shift (a << 2)  : %d << 2 = %d\n", a, a<<2)
	fmt.Printf("Right Shift (b >> 2) : %d >> 2 = %d\n", b, b>>2)
}
func main() {

	a := 15
	b := 10

	myXOR(a, b)
	myOR(a, b)
	myNOT(a, b)
	myAND(a, b)
	myShift(a, b)

}
