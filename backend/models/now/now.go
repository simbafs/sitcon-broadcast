package now

import "time"

type Now int64

var now Now = 0

func GetNow() Now {
	if now == 0 {
		return Now(time.Now().Unix())
	} else {
		return now
	}
}

func SetNow(t int64) {
	now = Now(t)
}

func ResetNow() {
	now = 0
}
