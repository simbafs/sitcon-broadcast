package now

import (
	"time"
)

var (
	zero = time.Time{}
	now  = time.Time{}
)

// no Create

func Read() time.Time {
	n := time.Now()
	if now.Equal(zero) {
		return n
	} else {
		return now
	}
}

func Update(t time.Time) {
	now = t
}

func Delete() {
	now = zero
}
