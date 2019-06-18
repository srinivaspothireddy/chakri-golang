package main

import "fmt"

//slices

// Slice allows you to group together VALUES of the same Type
func main() {
	//x:= type {values} //composite literal
	x := []int{4, 5, 6, 7, 8, 9}
	fmt.Println(x)
	fmt.Println(x[2])
	fmt.Println(len(x))
	fmt.Println(cap(x))
	for i := 0; i <= len(x); i++ {
		fmt.Println(i)
	}

	//Using for-range to display index and values
	for i, v := range x {
		fmt.Println(i, v)
	}

	// Slice using a Slice
	//x := []int{4, 5, 6, 7, 8, 9}
	fmt.Println(x[1])
	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[1:3]) //From index 1 upto index 3 (means it does not include index 3 value)

	//Append to a Slice
	//x := []int{4, 5, 6, 7, 8, 9}
	fmt.Println(x)
	x = append(x, 11, 12, 13, 14, 15)
	fmt.Println(x)

	y := []int{20, 30, 40}
	z := append(x, y...)

	fmt.Println(z)

	//deleting from a slice
	fmt.Println(x)
	x = append(x[:2], x[4:]...) // you are deleting the value in index 2,3 i.e. value = 6 and 7
	fmt.Println(x)

	//slice make
	m := make([]int, 10, 12)
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))
	m = append(m, 11)
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))

}
