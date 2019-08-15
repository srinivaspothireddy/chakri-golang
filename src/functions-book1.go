package main

import "fmt"

func main(){
	fmt.Println("Hey")
	fmt.Println(addition(2,3))
}

/*
Syntax:
-------
func name(paramaters-list) (results-list){
	body
}
 */

func addition(x int, _ int) int {
	return x
}
