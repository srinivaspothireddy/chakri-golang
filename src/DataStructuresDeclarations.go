package main

import "fmt"

func main() {
	//Array
	var arrayName1 []int
	fmt.Println(arrayName1)

	//Slices
	var sliceName [5]int
	fmt.Println(sliceName)

	sliceName1 := []int{1, 2, 3, 4, 5}
	fmt.Println(sliceName1)

	for i, val := range sliceName1 {
		fmt.Println(i, val)
	}

	sliceMake := make([]int, 10, 10)
	fmt.Println(sliceMake)

	//Map
	m1 := map[string]int{}
	fmt.Println(m1)

	m2 := map[string][]string{}
	fmt.Println(m2)

	//Struct
	var office struct {
		employees int
		chairs    string
	}
	fmt.Println(office)

}
