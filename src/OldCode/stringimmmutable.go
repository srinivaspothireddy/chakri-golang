package OldCode

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var x string = "abc"
	fmt.Println(x, &x, (*reflect.StringHeader)(unsafe.Pointer(&x)))
	x = "cde"
	fmt.Println(x, &x, (*reflect.StringHeader)(unsafe.Pointer(&x)))
	x = x[1:]
	fmt.Println(x, &x, (*reflect.StringHeader)(unsafe.Pointer(&x)))

}
