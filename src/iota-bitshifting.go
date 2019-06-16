package main

import "fmt"

var (
	a = 4
)

const (
	_ = iota
	d
	e
	f
)

const (
	//_ = iota
	//kb = 1024
	//mb = kb * 1024
	//gb = mb * 1024

	// The above can also be written as
	_ = iota
	//kb = 1024
	//kb = 1 << 10
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Println("Bit Shifting Code")
	fmt.Printf("%d\t\t%b\n", a, a)

	b := a << 1

	fmt.Printf("%d\t\t%b\n", b, b)

	//iota code
	fmt.Println("iota Code")
	//fmt.Printf("%d\t\t%b\n", c, c)
	fmt.Printf("%d\t\t%b\n", d, d)
	fmt.Printf("%d\t\t%b\n", e, e)
	fmt.Printf("%d\t\t%b\n", f, f)

	//Bytes Stuff
	fmt.Println("iota Bytes Code")
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}
