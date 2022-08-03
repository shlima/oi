package measure

import (
	"math"
	"time"
)

type Measure struct {
	start time.Time
}

func New() Measure {
	return Measure{start: time.Now()}
}

func (m Measure) Ms() int {
	got := time.Now().UnixMilli() - m.start.UnixMilli()
	if got > math.MaxInt {
		return math.MaxInt
	}

	return int(got)
}
