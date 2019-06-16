package main

import "fmt"

type hotdog int

var x3 hotdog

func main() {
	fmt.Println("x3: ", x3)
	fmt.Printf("%T\n", x3)
	x3 = 42
	fmt.Println("x3: ", x3)

}
