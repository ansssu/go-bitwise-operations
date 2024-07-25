package main

import "fmt"

func main() {
	// left shift
	x := 5
	fmt.Printf("Binary form of %d is %08b\n", x, x)
	y := x << 2 // shift left by 2 digits
	fmt.Printf("the value of %d << 2 is %d, wich in binary format is  %08b\n", x, y, y)

	//Right shift - unsigned
	z := uint8(24)
	fmt.Printf("Binary value of %d is %08b\n", z, z)
	w := z >> 3 // shift right by 3 digits
	fmt.Printf("the value of %d >> 3 is %d, wich in binary format is  %08b\n", z, w, w)

	//Right shift - signedd
	a := -8
	fmt.Printf("Binary value of %d is %08b\n", a, a)
	b := a >> 1 // shift right by 1 digit (arithmethic shift)
	fmt.Printf("the value of %d >> 1 is %d, wich in binary format is  %08b\n", a, b, b)
}
