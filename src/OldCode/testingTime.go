package OldCode

import (
	"fmt"
	"time"
)

func main() {
	// 1) First convert the strings to time Object
	// 2) Second, Format the time object to your wanted layout

	//fmt.Println("time.Now() : ", time.Now())
	//fmt.Println("time.Now().Unix() : ", time.Now().Unix())
	//fmt.Println("time.Now().UnixNano() : ", time.Now().UnixNano())
	//fmt.Println("time.Now().Format(\"2006-01-02T15:04:05Z\") : ", time.Now().Format("2006-01-02T15:04:05Z"))
	//fmt.Println("time.Now().Format(\"2006-01-02 15:04:05\") : ", time.Now().Format("2006-01-02 15:04:05"))
	//
	//t := time.Now()
	//fmt.Println(t.Format(time.ANSIC))

	//fmt.Println(formatISOTimestampHelper("2019-08-15T16:41:00Z"))

	// Use `time.Now` with `Unix` or `UnixNano` to get
	// elapsed time since the Unix epoch in seconds or
	// nanoseconds, respectively.
	now := time.Now()
	fmt.Println(now)

	secs := now.Unix()
	nanos := now.UnixNano()
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// You can also convert integer seconds or nanoseconds
	// since the epoch into the corresponding `time`.
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))

}

/*func formatISOTimestampHelper(isoTimestamp string) string {
	layoutISO := "2006-01-02 15:04:05"
	timeStampNoWhitespaces := strings.Join(strings.Fields(isoTimestamp), "")

	if !strings.ContainsAny(timeStampNoWhitespaces, "T & Z") {
		return isoTimestamp
	}

	t, err := time.Parse(time.RFC3339, isoTimestamp)
	if err != nil {
		return ""
	}
	return t.Format(layoutISO)
}*/
