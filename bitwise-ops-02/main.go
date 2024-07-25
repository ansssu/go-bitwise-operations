package main

import "fmt"

func main() {
	fmt.Printf("%d \t %08b\n", 1, 1)
	fmt.Printf("%d \t %08b\n", 1<<1, 1<<1)
	fmt.Printf("%d \t %08b\n", 1<<2, 1<<2)
	fmt.Printf("%d \t %08b\n", 1<<3, 1<<3)
	fmt.Printf("%d \t %08b\n", 1<<4, 1<<4)
	fmt.Printf("%d \t %08b\n", 1<<5, 1<<5)
	fmt.Printf("%d \t %08b\n", 1<<6, 1<<6)
	fmt.Printf("%d \t %08b\n", 1<<7, 1<<7)
	fmt.Printf("%d \t %08b\n", 1<<8, 1<<8)
	fmt.Printf("%d \t %09b\n", 1<<9, 1<<9)
	fmt.Printf("%d \t %010b\n", 1<<10, 1<<10)
}
