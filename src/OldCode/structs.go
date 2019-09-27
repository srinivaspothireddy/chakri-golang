package OldCode

import "fmt"

type person1 struct {
	firstname string
	lastname  string
	phno      int
}

type person2 struct {
	person1
	Address string
	ltk     bool
}

func main() {
	p1 := person1{
		firstname: "Srinivas",
		lastname:  "Pothireddy",
		phno:      8473003048,
	}

	fmt.Println(p1)
	fmt.Println(p1.firstname)
	fmt.Println(p1.lastname)
	fmt.Println(p1.phno)

	p2 := person2{
		ltk:     true,
		Address: "500 Park Way",
		person1: person1{
			firstname: "Srinivas",
			lastname:  "Pothireddy",
			phno:      8473003048,
		},
	}
	// Inner type (person1) gets promoted to the Outer Type p2.
	fmt.Println(p2.ltk, p2.Address, p2.firstname, p2.lastname, p2.phno)

	//Anonoymous Struct - Used to keep the code lean and clean. As in if you want to use that struct within only one area, then go for Anonymous.
	p3 := struct {
		name    string
		college string
		rollno  int
	}{
		name:    "James",
		college: "Northeastern",
		rollno:  32,
	}
	fmt.Println(p3)

}
