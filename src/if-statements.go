package main

import "fmt"

func main() {
	if true {
		fmt.Println("One")
	}

	if false {
		fmt.Println("Two")
	}

	if !true {
		fmt.Println("Three")
	}

	if !false {
		fmt.Println("Four")
	}

	//initialization statement
	if x7 := 42; x7 == 42 { // The scope of x here is only limited to this IF condition
		fmt.Println("001")
	}

	// If Else
	x := 40
	if x == 40 {
		fmt.Println("In 40")
	} else if x == 41 {
		fmt.Print("In 41")
	} else {
		fmt.Println("In ELse")
	}

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

}
