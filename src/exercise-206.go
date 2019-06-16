package main

import "fmt"

// Using iota print last 4 years

const(
	year = 2018
	year1 = year - iota
	year2 = year - iota
	year3 = year - iota
	year4 = year - iota

)
func main(){
	fmt.Println(year1, year2, year3, year4)
}
