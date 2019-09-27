package OldCode

import "fmt"

var total int = 0

func main() {
	i := 1988
	for i <= 2019 {
		fmt.Println(i)
		i++
		total++
	}
	fmt.Println("Total no of years lived : ", total, "years")

}
