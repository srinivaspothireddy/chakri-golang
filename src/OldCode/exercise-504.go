package OldCode

import "fmt"

func main() {

	a1 := struct {
		eyes  int
		skin  string
		bones bool
	}{
		eyes:  2,
		skin:  "black",
		bones: true,
	}
	fmt.Println(a1)
}
