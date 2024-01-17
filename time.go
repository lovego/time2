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
	timeFormat = "2006-01-02 15:04:05"
)

func Now() Time {
	return Time{time.Now()}
}

func NowPtr() *Time {
	return &Time{time.Now()}
}

func New(year, month, day, hour, minute, second int) Time {
	return Time{time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)}
}

func Parse(str string) (Time, error) {
	str = strings.Trim(str, "\"")
	if str == "" || str == "null" {
		return Time{}, nil
	}

	t, err := time.Parse(timeFormat, str)
	if err != nil {
		return Time{}, err
	}

	return Time{t}, nil
}

func (t Time) String() string {
	if t.Time.IsZero() {
		return ""
	}
	return t.Format(timeFormat)
}

func (t Time) MarshalJSON() ([]byte, error) {
	if t.Time.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + t.Format(timeFormat) + `"`), nil
}

func (t *Time) UnmarshalJSON(b []byte) error {
	d, err := Parse(string(b))
	if err != nil {
		return err
	}

	*t = d
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
		*t = Time{v}
		return nil
	}
	return fmt.Errorf("can not convert %v to t.Time", value)
}
