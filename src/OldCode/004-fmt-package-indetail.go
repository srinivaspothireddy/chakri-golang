package OldCode

import "fmt"

var y1 = 42

func main() {

	fmt.Println(y1)
	fmt.Printf("%T\n", y1)
	fmt.Printf("%b\n", y1)
	fmt.Printf("%x\n", y1)
	fmt.Printf("%#x\n", y1)

	s := fmt.Sprintf("%x", y1) //Prints the value assining to a String, "s" here
	fmt.Println(s)

	//	fmt.Fprintf("%#x\n",y1) // prints to a file or web server

}
