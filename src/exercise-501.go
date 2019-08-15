package main

import "fmt"

type myperson struct {
	fn               string
	ln               string
	icecreamflavours []string
}

func main() {

	me1 := myperson{
		fn:               "Srinivas",
		ln:               "Pothireddy",
		icecreamflavours: []string{"Vanilla", "Strawberry", "Chocolate"},
	}

	me2 := myperson{
		fn:               "Srinivas",
		ln:               "Pothireddy",
		icecreamflavours: []string{"Vanilla", "Strawberry", "Chocolate"},
	}

	fmt.Println(me1)
	fmt.Println(me2)

	for _, v := range me1.icecreamflavours {
		fmt.Println(v)

	}

}
