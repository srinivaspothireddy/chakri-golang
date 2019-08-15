package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println("UnixTime :- ", formatToEpochTimestamp("2019-08-13 15:04:05"))
	//fmt.Println(formatUnixTimeHelper("2006-01-02T15:04:05Z", 1565889378.0000))

	fmt.Println("time.Now() : ", time.Now())
	fmt.Println("time.Now().Unix() : ", time.Now().Unix())
	fmt.Println("time.Now().UnixNano() : ", time.Now().UnixNano())
	fmt.Println("time.Now().Format(\"2006-01-02T15:04:05Z\") : ", time.Now().Format("2006-01-02T15:04:05Z"))
	fmt.Println("time.Now().Format(\"2006-01-02 15:04:05\") : ", time.Now().Format("2006-01-02 15:04:05"))


	t := time.Now()
	fmt.Println(t.Format(time.ANSIC))
}
