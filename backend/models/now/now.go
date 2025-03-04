package now

import (
	"time"
)

type Now struct {
	time.Time
}

func (n Now) MarshalJSON() ([]byte, error) {
	loc, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return nil, err
	}

	return []byte(`"` + n.In(loc).Format("2006-01-02 15:04:05") + `"`), nil
}

var (
	zero = Now{}
	now  = Now{}
)

// no Create

var tw = time.FixedZone("Asia/Taipei", 8)

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
