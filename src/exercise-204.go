package main

import "fmt"

// Print a numer in binary, decimal and hexa decimal and now shift it one bit left and print the same that new value.

func main() {
	e4 := 8
	fmt.Printf("%d\t%b\t%#X\n", e4, e4, e4)
	e41 := e4 << 1
	fmt.Printf("%d\t%b\t%#X", e41, e41, e41)

}
