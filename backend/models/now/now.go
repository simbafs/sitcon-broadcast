package now

import (
	"time"
)

var (
	zero = time.Time{}
	now  = time.Time{}
)

// no Create

var tw = time.FixedZone("Asia/Taipei", 8)

func Read() time.Time {
	n := time.Now()
	if now.Equal(zero) {
		return n.In(tw)
	} else {
		return now.In(tw)
	}
}

func Update(t time.Time) {
	now = t
}

func Delete() {
	now = zero
}
