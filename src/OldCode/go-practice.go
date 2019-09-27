package OldCode

import (
	"fmt"
)

var y21 int

type persontype struct {
	fname string
	lname string
}

type secretAgent struct {
	persontype
	licenseToKill bool
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func (p persontype) speak() {
	fmt.Println(p.fname, `says, "Good Morning!!!"`)

}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Good Morning!!!"`)

}

func main() {
	x := 5
	fmt.Println(x)
	fmt.Printf("%T", x)

	y21 = 7
	fmt.Println(y21)
	fmt.Printf("%T", y21)

	//Slice of int
	xi := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(xi)

	//map
	m := map[string]int{
		"Srini":  5,
		"Chakri": 7,
	}
	fmt.Println(m)

	//struct
	p1 := persontype{
		"Miss",
		"MoneyPenny",
	}
	fmt.Println(p1)

	//Function
	p1.speak()

	sa1 := secretAgent{
		persontype{
			"James",
			"Bond",
		},
		true,
	}
	fmt.Println(sa1)

	///composition
	sa1.speak()
	sa1.persontype.speak()

	//Interface & Polymosphism
	saySomething(p1)
	saySomething(sa1)

}
