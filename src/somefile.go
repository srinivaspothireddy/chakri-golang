package main

import "fmt"

type vehicle1 struct {
	doors int
	color string
}

type truck1 struct {
	vehicle1
	fourwheel bool
}

type sedan1 struct {
	vehicle1
	luxury bool
}

func main() {
	t1 := truck1{
		vehicle1: vehicle1{
			doors: 4,
			color: "blue",
		},
		fourwheel: false,
	}
	fmt.Println(t1)

	s1 := sedan1{
		vehicle1: vehicle1{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(s1)

}
