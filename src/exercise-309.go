package main

import "fmt"

var favsport string = "surfing"

func main() {
	switch favsport {
	case "surfing":
		fmt.Println("Yo, this is true")
		fallthrough
	case "something":
		fmt.Println("Yo this is false")
		fallthrough
	default:
		fmt.Println("Yo this is in default")
	}
}
