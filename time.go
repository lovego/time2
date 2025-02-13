package time2

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

const (
	timeFormat = time.RFC3339
	format     = "2006-01-02 15:04:05"
)

func Now() Time {
	return Time{time.Now()}
}

func NowPtr() *Time {
	return &Time{time.Now()}
}
func NewPtr(t time.Time) *Time {
	return &Time{t}
}

func New(year, month, day, hour, minute, second int) Time {
	return Time{time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)}
}

func Parse(str string) (Time, error) {
	str = strings.Trim(str, "\"")
	if str == "" || str == "null" {
		return Time{}, nil
	}
	// 加载目标时区
	loc := time.Local
	if GetTimeZone != nil {
		l := GetTimeZone()
		if l != nil {
			loc = l
		}
	}
	t, err := time.ParseInLocation(timeFormat, str, loc)
	if err != nil {
		t, err = time.ParseInLocation(format, str, loc)
		if err != nil {
			return Time{}, err
		}
	}

	return Time{t}, nil
}

func (t Time) String() string {
	if t.Time.IsZero() {
		return ""
	}

	return t.Format(format)
}

func (t Time) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte(`""`), nil
	}

	// 加载目标时区
	if GetTimeZone != nil {
		loc := GetTimeZone()
		if loc == nil {
			return []byte(`"` + t.Format(format) + `"`), nil
		}
		t.Time = t.Time.In(loc)
	}

	return []byte(`"` + t.Format(format) + `"`), nil
}

//12:00 +8  12:00 +0 ==> 12-8
//4:00 +0
//10:00  +6

func (t *Time) UnmarshalJSON(b []byte) error {
	d, err := Parse(string(b))
	if err != nil {
		return err
	}
	*t = d
	t.Time = t.Time.Local()
	return nil
}

func (t Time) Value() (driver.Value, error) {
	if t.Time.IsZero() {
		return []byte("NULL"), nil
	}
	return []byte(t.Format("'2006-01-02T15:04:05.999999Z07:00'")), nil
}

func (t *Time) Scan(value interface{}) error {
	if value == nil {
		*t = Time{}
		return nil
	}

	v, ok := value.(time.Time)
	if ok {
		*t = Time{v.Local()}
		return nil
	}
	return fmt.Errorf("can not convert %v to t.Time", value)
}

var GetTimeZone = func() *time.Location {
	// todo 获取自定义时区,调整该函数

	//fmt.Println("获取自定义时区,调整该函数")

	return time.Local
}

func (t Time) Add(d time.Duration) Time {
	return Time{t.Time.Add(d)}
}

func (t Time) Sub(tt Time) time.Duration {
	return t.Time.Sub(tt.Time)
}

func (t Time) After(tt Time) bool {
	return t.Time.After(tt.Time)
}

func (t Time) Before(tt Time) bool {
	return t.Time.Before(tt.Time)
}
func (t Time) AddDate(years int, months int, days int) Time {
	return Time{t.Time.AddDate(years, months, days)}
}
func Since(tt Time) time.Duration {
	return time.Since(tt.Time)
}
