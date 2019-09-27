package OldCode

import "fmt"

type myperson1 struct {
	fn               string
	ln               string
	icecreamflavours []string
}

func main() {

	me1 := myperson1{
		fn:               "Srinivas",
		ln:               "Pothireddy",
		icecreamflavours: []string{"Vanilla", "Strawberry", "Chocolate"},
	}

	me2 := myperson1{
		fn:               "James",
		ln:               "Bond",
		icecreamflavours: []string{"Blackcurrent", "Orange", "Buuterscotch"},
	}

	fmt.Println(me1)
	fmt.Println(me2)

	//MAP
	map1 := map[string][]string{
		me1.ln: me1.icecreamflavours,
		me2.ln: me2.icecreamflavours,
	}
	fmt.Println(map1)

	for k, v := range map1 {
		fmt.Println("Printing flavors for : ", k)
		fmt.Println(v)

	}

}
