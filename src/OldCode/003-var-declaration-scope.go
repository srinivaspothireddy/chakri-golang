package OldCode

import (
	"fmt"
)

var y = 100 // Scope is with in this while package

var z int // By default assigns 0 to z
//Each element of such a variable or value is set to the zero value for its type: false for booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps.

func main() {
	// DECLARE A VARIABLE AND ASSIGN A VALUE
	x := 10 // Scope is limited to with in this function
	fmt.Println(x)

	//ASSIGNING A VALUE
	x = 45 // Scope is limited to with in this function
	fmt.Println(x)

	fmt.Println(y)
	foo1()
	fmt.Println(z)
}

func foo1() {
	fmt.Println(y)

}
