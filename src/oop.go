package main

import "fmt"

// we create values of ccertain type that are stored in variables and those variables have identifiers

var no int

type person struct {
	first string
	last  string
}

func main() {

	fmt.Println(no)
	p1 := person{
		first: "James",
		last:  "Bond",
	}

	fmt.Println(p1)

}
