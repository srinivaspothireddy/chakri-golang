package functions

import (
	"fmt"
	"regexp"
	"time"
)

/*
Returning a function from a function
Callback - Is passing a func as an argument
*/
var fact int

func main() {
	//s1 := fooNamee()
	//fmt.Println(s1)
	//
	//b1 := booNamee()
	//fmt.Println(b1())
	//
	//ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//sum := sum1(ii...)
	//fmt.Println("The sum all numbers :", sum)
	//
	//s2 := even(sum1, ii...)
	//fmt.Println("Sum all of all even numbers", s2)
	//
	//s3 := odd(sum1, ii...)
	//fmt.Println("Sum all of all odd numbers", s3)

	//fff := factorial1(5)
	//fmt.Println(fff)

	layoutNonISO := "2006-01-02 15:04:05"
	//inputDateTime := "2019-07-28T03:00:00Z"
	////removeTZ(inputDateTime)
	//
	//t, _ := time.Parse(time.UnixDate, inputDateTime)
	//fmt.Println(t)
	//fmt.Println(t.Format(layoutNonISO))

	p := fmt.Println
	//t := time.Now()
	//p(t)
	//p(t.Format(layoutNonISO))

	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41Z")
	p(t1)
	p(t1.Format(layoutNonISO))

}

func removeTZ(inputDateTime string) {
	T := regexp.MustCompile(`T`)
	fmt.Println(T.Split(inputDateTime, -1))

}

func fooNamee() string {
	s := "Hello World"
	return s // returning a String
}

func booNamee() func() int {
	return func() int {
		return 55
	}
}

/*
Callback - A callback is passing a function as an argument
*/
func sum1(xi ...int) int {

	sum := 0
	for _, v := range xi {
		sum = sum + v
	}
	return sum
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

//recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}

}

func factorial1(n int) int {
	for i := 1; i <= n; i++ {
		n = n*n - i
	}

	return n

}
