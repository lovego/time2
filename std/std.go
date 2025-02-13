package std

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

//type Time struct {
//	time.Time
//}

type (
	Time     = time.Time
	Duration = time.Duration
	Location = time.Location
	Month    = time.Month
	Weekday  = time.Weekday
	Timer    = time.Timer
	Ticker   = time.Ticker
)

// 重导出 time 包的常量
const (
	January    = time.January
	February   = time.February
	Monday     = time.Monday
	Second     = time.Second
	Nanosecond = time.Nanosecond
)

// 重导出 time 包的变量
var (
	UTC   = time.UTC
	Local = time.Local
)

// 重导出 time 包的函数
var (
	Now   = time.Now
	Since = time.Since
	Until = time.Until
	Date  = time.Date
	//Parse         = time.Parse
	Sleep         = time.Sleep
	NewTicker     = time.NewTicker
	NewTimer      = time.NewTimer
	After         = time.After
	AfterFunc     = time.AfterFunc
	Tick          = time.Tick
	ParseDuration = time.ParseDuration
)

const (
	timeFormat = "2006-01-02 15:04:05"
)

//func Now() Time {
//	return Time{time.Now()}
//}

func NowPtr() *Time {
	now := time.Now()
	return &now
}

func New(year, month, day, hour, minute, second int) Time {
	return time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
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

	return t, nil
}

func (t Time) String() string {
	if t.IsZero() {
		return ""
	}

	return t.Local().Format(timeFormat)
}

func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + t.Format(timeFormat) + `"`), nil
}

func (t *Time) UnmarshalJSON(b []byte) error {
	d, err := Parse(string(b))
	if err != nil {
		return err
	}
	*t = d
	// 加载目标时区
	loc := GetTimeZone()
	if loc == nil {
		return nil
	}

	// 转换到目标时区
	t.Time = t.Time.In(loc)
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

	fmt.Println("获取自定义时区,调整该函数")

	return time.UTC
}
