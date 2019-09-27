package OldCode

import "fmt"

var x = 5

func main() {
	switch x {
	case 5:
		fmt.Println("Yes the value is 5")
		fallthrough
	case 7:
		fmt.Println("No The value is not 5")
	default:
		fmt.Println("In default")
	}
}
