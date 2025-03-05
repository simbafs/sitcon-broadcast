package now

import (
	"time"
)

type Now struct {
	time.Time
}

func (n Now) MarshalJSON() ([]byte, error) {
	return []byte(`"` + n.Format("2006-01-02T15:04:05.000000000Z") + `"`), nil
}

var (
	zero = Now{}
	now  = Now{}
)

// no Create

func getRealNow() Now {
	return Now{time.Now()}
}

func Read() Now {
	n := getRealNow()
	if now.Equal(zero.Time) {
		return n
	} else {
		return now
	}
}

func Update(t Now) {
	now = t
}

func Delete() {
	now = zero
}
