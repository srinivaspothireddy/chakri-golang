package OldCode

import "fmt"

// Different FOR loop formats
func main() {

	// FOR Clause
	for i := 1; i <= 25; i++ {
		fmt.Println("i: ", i)
		for j := 1; j <= 10; j++ {
			fmt.Println("j: ", j)
		}

	}

	// FOR condition
	fmt.Println("For Condition")
	rollno := 5
	for rollno <= 15 {
		fmt.Println("Yes, The rollno value is 5.")
		rollno++
	}

	//For Range
	fmt.Println("For Range")

	//Continue

	//break

}
