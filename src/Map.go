package main

import "fmt"

//Maps are key value stores
func main() {
	m := map[string]int{
		"James":           32,
		"Miss MoneyPenny": 27,
		"MaheshBabu":      45,
	}

	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Charkri"]) //This Key is not declared, But it prints the value as ZERO by default.

	//In order to check this, we can declare as below.
	// ok will be false if there is no key declared
	//Thus it will check and Prints the value only if ok is true
	if v, ok := m["chakerl"]; ok {
		fmt.Println(v)
	}

	// Map add element and Range
	//Adding an element to a map
	m["newName"] = 33
	fmt.Println(m)

	//Range
	for k, v := range m {
		fmt.Println(k, v)
	}

	//delete
	delete(m, "newName")
	fmt.Println(m)

	if v, ok := m["adfdf"]; ok {
		fmt.Println(v)
		delete(m, "adfdf")

	}

}
