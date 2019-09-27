package OldCode

import "fmt"

func main() {
	switch {
	case true:
		fmt.Println("Yo, this is true")
		fallthrough
	case false:
		fmt.Println("Yo this is false")
		fallthrough
	default:
		fmt.Println("Yo this is in default")
	}
}
