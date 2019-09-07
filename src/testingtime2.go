package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	returnedFormattedTimestamp := formatISOWithNoTZHelper("2019-08-18T17:40:00Z0700")
	fmt.Println("The converted timestamp without designators is : ", returnedFormattedTimestamp)

	returnedUnixTimestamp := formatToEpochTimestamp("2006-01-02T15:04:05.000Z", "2017-06-29T15:45:33.345Z")
	fmt.Println("The converted Unix Timestamp is :", returnedUnixTimestamp)

}

func formatISOWithNoTZHelper(inputTimestamp string) string {
	layoutISOWithNoTZ := "2006-01-02 15:04:05"
	if len(inputTimestamp) == 20 && strings.Index(inputTimestamp, "T") == 10 && strings.Index(inputTimestamp, "Z") == 19 {
		t, err := time.Parse(time.RFC3339, inputTimestamp)
		if err != nil {
			return ""
		}
		return t.Format(layoutISOWithNoTZ)
	} else if len(inputTimestamp) == 19 && strings.Index(inputTimestamp, "T") == -1 && strings.Index(inputTimestamp, "Z") == -1 {
		return inputTimestamp
	} else {
		return ""
	}
}

func formatToEpochTimestamp(formatTimestamp string, inputTimestamp string) float64 {

	t, err := time.Parse(formatTimestamp, inputTimestamp)
	if err != nil {
		return 0
	}

	unixTSWithFractionSecs := float64(t.UnixNano()) / 1000000000
	return unixTSWithFractionSecs

}
