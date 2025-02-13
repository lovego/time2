package time2

import (
	"encoding/json"
	"fmt"
	"time"
)

func ExampleParse() {
	//req.Body: parsing time ""2025-01-10 15:24:58"" as ""2006-01-02T15:04:05Z07:00"": cannot parse " 15:24:58"" as "T"
	fmt.Println(Parse("2019-09-12 12:01:02"))
	// Output: 2019-09-12 12:01:02 <nil>
}

func ExampleParse2() {
	fmt.Println(Parse("2019-09-12T12:01:02.971689452+08:00"))
	// Output: 2019-09-12 12:01:02 <nil>
}

func ExampleMarshalJSON() {
	var t = Time{}
	b, err := json.Marshal(t)
	fmt.Println(string(b), err)

	t = New(2019, 9, 12, 12, 1, 2)
	b, err = json.Marshal(t)
	fmt.Println(string(b), err)
	// Output:
	// "" <nil>
	// "2019-09-12 12:01:02" <nil>
}

func ExampleUnmarshalJSON() {
	var t = Time{}
	err := json.Unmarshal([]byte(`"2019-09-12 12:01:02"`), &t)
	fmt.Println(t, err)

	t = Time{}
	err = json.Unmarshal([]byte(`2019-09-12 12:01:02`), &t)
	fmt.Println(t, err)

	// Output:
	// 2019-09-12 20:01:02 <nil>
	//  invalid character '-' after top-level value
}

func ExampleValue() {
	b, err := Time{}.Value()
	fmt.Println(string(b.([]byte)), err)

	t := Time{time.Date(2019, 9, 12, 12, 1, 2, 0, time.UTC)}
	b, err = t.Value()
	fmt.Println(string(b.([]byte)), err)

	// Output:
	// NULL <nil>
	// '2019-09-12T12:01:02Z' <nil>
}
