package now

import (
	"time"
)

type Time int

func (t Time) Hour() int {
	return int(t) / 60
}

func (t Time) Minute() int {
	return int(t) % 60
}

func (t Time) Time() time.Time {
	return time.Date(2024, 3, 9, t.Hour(), t.Minute(), 0, 0, time.FixedZone("Asia/Taipei", 8))
}

var now = Time(0)

func GetNow() Time {
	n := time.Now()
	if now == 0 {
		return Time(n.Hour()*60 + n.Minute())
	} else {
		return now
	}
}

func SetNow(t int) {
	now = Time(t)
}

func ClearNow() {
	now = 0
}
