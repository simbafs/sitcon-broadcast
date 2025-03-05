package now

import (
	"time"
)

var tw = time.FixedZone("Asia/Taipei", 8)

type Now struct {
	time.Time
}

func (n Now) MarshalJSON() ([]byte, error) {
	return []byte(`"` + n.In(tw).Format("2006-01-02 15:04:05") + `"`), nil
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
