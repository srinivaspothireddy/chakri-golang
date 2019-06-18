package main

import "fmt"

func main() {
	slice1 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(slice1)
	slice1 = append(slice1, 52)
	fmt.Println(slice1)

	slice1 = append(slice1, 53, 54, 55)
	fmt.Println(slice1)

	slice2 := []int{56, 57, 58, 59, 60}
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)

}
