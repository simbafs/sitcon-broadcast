package now

import (
	"time"

	"backend/sse"
)

type Now int64

var now Now = 0

func GetNow() Now {
	if now == 0 {
		return Now(time.Now().Unix())
	} else {
		return now
	}
}

func SetNow(t int64, send chan sse.Msg) {
	now = Now(t)
	send <- sse.Msg{
		Topic: []string{"now"},
		Data:  now,
	}
}

func ResetNow(send chan sse.Msg) {
	now = 0
	send <- sse.Msg{
		Topic: []string{"now"},
		Data:  now,
	}
}
