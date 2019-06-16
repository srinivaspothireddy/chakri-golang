package main

import "fmt"

func main() {
	fmt.Println("Hello World! This is my main first program...")
	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()

	n, err := fmt.Println("I am exiting the main function")
	fmt.Println(n)
	fmt.Println(err)

}

func foo() {
	fmt.Println("Yoo.. I am in foo function.")
}

func bar() {
	fmt.Println("Yo.. I am in function bar.")
}
