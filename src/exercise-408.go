package main

import (
	"fmt"
)

func main() {
	m1 := map[string][]string{
		"Bond_james":      []string{"Shaken, not Stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being Evil", "Ice cream", "Sunsets"},
	}

	fmt.Println(m1)
	for k, v := range m1 {
		//fmt.Println(k,v)
		for i, val := range v {
			fmt.Println(k, i, val)
		}
	}
}
