package functions

import (
	"fmt"
)

//declare a int variable
var someinteger int

/*
METHOD
*/
//Struct
type personStruct struct {
	first string
	last  string
}

type secretAgentStruct struct {
	personStruct
	ltk bool
}

//creating a Method
func (s secretAgentStruct) speak() { // (s secretAgentStruct) - is the receiver for function speak()
	fmt.Println("I am ", s.first, s.last)
}

// Breaks the code into modular
func main() {

	fooName()
	fooName1("Srini")
	fooName2("Srini", 30)
	s1 := fooName3("Chakravarthy")
	fmt.Print(s1)

	defer wooName() // Always runs at the end of the declared function. In this case at the end of main() function.

	s2, x := fooName4("Chakravarthy", 29)
	fmt.Println("This is fooName4 : ", s2, x)

	a, x := fooName5("firstname", "lastname", 35)
	fmt.Println(a, x)

	// Variadic Parameters
	//sum := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	//fmt.Println("The sum is: ", sum)

	//The above can sum function with variadic paramaters can also be declared as below, by Unfurling a Slice
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum1 := sum(xi...) // This is called unfurling a slice to send as a variadic parameter
	fmt.Println("The sum is: ", sum1)

	//Or you can also pass nothing. i.e.
	sum2 := sum()
	fmt.Println("The sum is :", sum2)

	wooName1()

	//Belongs to Method
	sa1 := secretAgentStruct{
		personStruct: personStruct{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgentStruct{
		personStruct: personStruct{
			"Money",
			"Penny",
		},
		ltk: false,
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	sa1.speak()
	sa2.speak()

}

//func (r receiver) identifier (parameters) (returns(s)) { code ..... }
/*
func fooName4(s1 string, x int) (string,int) {
	fmt.Println("This is ",s1," Guys")
	return s1,x
}
From the above,
fooName4 = identifier
(s1 string, x int) = Parameters
(string,int) = returns

*/

func fooName() {
	fmt.Println("hello from foo")
}

func fooName1(s string) {
	fmt.Println("hello from fooName1", s)
}

func fooName2(s string, x int) {
	fmt.Println("hello from fooName1", s, "Age", x)
}

//everything in GO is PASS BY VALUE

// Return a String
func fooName3(s string) string {
	fmt.Println("hello from fooName1", s)
	return s
}

// Multiple returns

func fooName4(s1 string, x int) (string, int) {
	fmt.Println("This is ", s1, " Guys")
	return s1, x
}

func fooName5(s1, s2 string, x int) (string, int) {
	a := fmt.Sprint(s1, " ", s2, "says hello to", x, "people.")
	return a, x
}

// Variadic Parameters
/*
Here, the input parameters are unlimited.
Stored as a Slice
... is an operator similar to && , || etc..
*/
func sum(y ...int) int { // ...int means zero to many paramters. You can also pass 0 parameters
	// This always has to be the final paramamter.
	// Example sum(s string, y ...int)

	fmt.Println("In booName")
	fmt.Println(y)
	fmt.Printf("%T \n", y) //Those values are being

	sum := 0
	for i, v := range y {
		fmt.Println(i, v)
		sum = sum + y[i]

	}
	//fmt.Println("The sum is: ",sum)
	return sum
}

/*
defer - is a keyword
*/
func wooName() {
	fmt.Println("I am in wooName - I am DEFER")
}

func wooName1() {
	fmt.Println("I am in wooName1")
}

/*
Methods -
*/
func kooName() {
	//fmt.Println(sa1)
}
