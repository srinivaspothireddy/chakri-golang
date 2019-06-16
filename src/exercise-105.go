package main

import "fmt"

type hotdog1 int

var x4 hotdog1
var y4 int

func main() {
	fmt.Println("x4: ", x4)
	fmt.Printf("%T\n", x4)
	x4 = 42
	fmt.Println("x4: ", x4)

	y4 = int(x4)
	fmt.Println("y4: ", y4)
	fmt.Printf("%T\n", y4)
}
