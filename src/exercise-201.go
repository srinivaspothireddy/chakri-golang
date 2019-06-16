package main

import "fmt"

// Print the number is decimal, binary and hexadecimal format

func main() {
	num1 := 6
	fmt.Printf("The decimal format is : %d\n", num1)
	fmt.Printf("The binary format is : %b\n", num1)
	fmt.Printf("The Hexadecimal format is : %#X\n", num1)
}
