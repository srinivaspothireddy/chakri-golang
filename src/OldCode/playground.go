package OldCode

import "fmt"

var array1 [5]int
var slicevar []int

var onevar int
var newarr []string

var map1 = map[string]int{
	"James":           32,
	"Miss MoneyPenny": 27,
	"Mahesh Babu":     45,
}

type person3 struct {
	firstname string
	lastname  string
	rollno    int
}

type person4 struct {
	person3
	Id int
}

func main() {

	p3 := person3{
		firstname: "Srinivas",
		lastname:  "Pothireddy",
		rollno:    32,
	}

	fmt.Println(p3.firstname)

	p4 := person4{
		person3: person3{
			firstname: "Srinivas",
			lastname:  "Pothireddy",
			rollno:    32,
		},
		Id: 17653,
	}

	fmt.Println(p4)

	fmt.Println(newarr)
	fmt.Print(onevar)

	fmt.Println(array1)
	array1[0] = 1
	fmt.Println(array1)

	slicevar = []int{1, 2, 3, 4, 5, 6}

	fmt.Println(slicevar)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice1)

	fmt.Println(slice1[:2])
	fmt.Println(slice1[2:])
	fmt.Println(slice1[2:6])

	fmt.Println("In For Loop")

	for i := 0; i < len(slice1); i++ {
		fmt.Println(slice1[i])
	}

	fmt.Println("Using Range")
	for i, v := range slice1 {
		fmt.Println(i, v)
	}

	slice2 := append(slice1, 11, 12, 13)
	fmt.Println(slice2)

	slice3 := append(slice1, slice2...)
	fmt.Println(slice3)

	slice4 := append(slice1[4:7], slice3[4:7]...)
	fmt.Println(slice4)

	slice5 := make([]int, 5, 10)
	slice5 = []int{11, 22, 33, 44, 55}
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))

	slice5 = append(slice5, 66)

	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))

	//fmt.Println(len(slice5))
	//fmt.Println(cap(slice5))

	map1 := map[string]int{
		"James":           32,
		"Miss MoneyPenny": 27,
		"Mahesh Babu":     45,
	}
	fmt.Println(map1)
	for k, v := range map1 {
		fmt.Println(k, v)
	}
	fmt.Println(map1["James"])
	fmt.Println(map1["James1"])
	map1["James1"] = 30
	fmt.Println(map1["James1"])
	fmt.Println(map1)

	if v, ok := map1["James1"]; ok {
		fmt.Println(v)
	}

	delete(map1, "James1")

	fmt.Println(map1)

	map2 := map[string][]string{
		"Srini": {"one", "two", "three"},
	}
	fmt.Println(map2)

}
