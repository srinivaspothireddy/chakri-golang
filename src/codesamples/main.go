package main

import "fmt"

var t int = 5
type hotdog int
var u hotdog


func main() {
	fmt.Println("This is my first code")
	foo()
	fmt.Println("Back to Main")

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even : ", i)
		}
	}


	fmt.Println(t)
	fmt.Printf("%T\n", t)

	fmt.Println(u)
	fmt.Printf("%T\n", u)

	t = int(u)

	fmt.Println("Loop completed")




}

func foo() {
	fmt.Println("I am in foo")
}
