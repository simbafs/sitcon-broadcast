package now

import (
	"time"
)

// Now is unix timestamp in seconds
type Now int64

var now Now = 0

// no Create

func getRealNow() Now {
	return Now(time.Now().Unix())
}

func Read() Now {
	n := getRealNow()
	if now == 0 {
		return n
	} else {
		return now
	}
}

func Update(t Now) {
	now = t
}

func Delete() {
	now = 0
}
