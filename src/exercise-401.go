package main

import "fmt"

func main() {
	array1 := [5]int{1, 2, 3, 4, 5}

	for _, v := range array1 {
		fmt.Println(v)
	}
	fmt.Printf("%T", array1)
}
