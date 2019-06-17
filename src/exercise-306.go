package main

import "fmt"

func main() {
	if true {
		fmt.Println("Hey it is true")
	} else if 3 == 3 {
		fmt.Println("This is also true.")
	} else {
		fmt.Println("This is in else")
	}
}
