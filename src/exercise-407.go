package main

import "fmt"

func main() {
	mslice1 := []string{"James", "Bond", "Shaken", "not stirred"}
	fmt.Println(mslice1)
	mslice2 := []string{"Miss", "Moneypenny", "Helloooooo", "James"}
	fmt.Println(mslice2)

	mslice3 := [][]string{mslice1, mslice2}
	fmt.Println(mslice3)

	for i, v := range mslice3 {
		fmt.Println(i, v)
		for i1, v1 := range mslice3[i] {
			fmt.Println(i1, v1)
		}
	}

}
