package OldCode

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(slice1)
	fmt.Printf("%T\n", slice1)

	for i, v := range slice1 {
		fmt.Println(i, v)

	}
}
