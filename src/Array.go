package main

import "fmt"

//Data Structures
func main() {
	//Array - To Group same type of values

	var x [5]int //You need to specify the size and type
	fmt.Println(x)
	x[0] = 123
	x[1] = 234
	x[2] = 345
	x[3] = 456
	x[4] = 567
	for i := 0; i <= 4; i++ {
		fmt.Println(x[i])
	}

}
