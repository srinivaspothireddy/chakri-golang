package functions

import "fmt"

//Format to declare anything -> "<Keyword> <Identifier> <Type>"

/*
Interfaces and Polymorphism

Interfaces allows to define the behaviour of a Function
Also allows to do the Polymorphism

*/
//An empty interface "interface{}" means any type of values will be accepted.
//A VALUE can be of more than one type.

type person5 struct {
	firstName string
	lastName  string
}

func (p person5) speak() {
	fmt.Println("I am ", p.firstName, p.lastName)
}

type secretAgent struct {
	person5
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am ", s.firstName, s.lastName, "and my License to Kill is : ", s.ltk)
}

//Interface (If you have this method "speak()", then you are also of my type "human")
type human interface {
	speak()
}

//Polymorphism
/*
func barName(h human) {
	fmt.Println("I am in bar name", h)
}
*/

//Assertion of Types
func barName(h human) {
	switch h.(type) {
	case person5:
		fmt.Println("I am in bar name", h.(person5).firstName)
	case secretAgent:
		fmt.Println("I am in bar name", h.(secretAgent).firstName)
	}
	fmt.Println("I am in bar name", h)
}

type burger int

func main() {

	pe1 := person5{
		"Dr.",
		"Strange",
	}

	sa1 := secretAgent{
		person5: person5{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person5: person5{
			"Money",
			"Penny",
		},
		ltk: false,
	}

	pe1.speak()

	sa1.speak()
	sa2.speak()

	barName(pe1)
	barName(sa1)
	barName(sa2)

	// conversion
	var x2 burger = 42
	fmt.Println(x2)
	fmt.Printf("%T \n", x2)

	var y2 int = 52
	fmt.Println(y2)
	fmt.Printf("%T \n", y2)

	y2 = int(x2) // Type conversion. Converting burger type into int type
	fmt.Println(y2)
	fmt.Printf("%T \n", y2)

	anonymous() // Non-Anonymous function, because it has a name

	//Anonymous function with no arguments
	func() {
		fmt.Println("In am in an anonymous function")
	}()

	//Anonymous function with arguments
	func(x int) {
		fmt.Println("In am in an anonymous function", x)
	}(42)

	/*
		Func Expression
	*/
	f := func() {
		fmt.Println("This is my first func expression")
	}
	f()

	f1 := func(x int) {
		fmt.Println("This is my first func expression", x)
	}
	f1(43)
}

/*
Anonymous Function
*/
func anonymous() {
	fmt.Println("I am in non - anonymous")
}
