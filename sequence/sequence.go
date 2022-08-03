package sequence

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// Sequence используется в тестах, для генерации различных последовательностей
type Sequence struct {
	int64  int64
	int32  int32
	source rand.Source
	rand   *rand.Rand
}

// New вернет новый Sequence
func New(seed int64) *Sequence {
	source := rand.NewSource(seed)

	return &Sequence{
		source: source,
		rand:   rand.New(source),
	}
}

func (s *Sequence) NextInt64() int64 {
	return atomic.AddInt64(&s.int64, 1)
}

func (s *Sequence) NextInt32() int32 {
	return atomic.AddInt32(&s.int32, 1)
}

func (s *Sequence) NextString() string {
	return fmt.Sprintf("%d", atomic.AddInt64(&s.int64, 1))
}

// RandomInt64 may be used to generate random amount of money
func (s *Sequence) RandomInt64() int64 {
	return s.rand.Int63()
}

// RandomInt32 may be used to generate random amount of money
func (s *Sequence) RandomInt32() int32 {
	return s.rand.Int31()
}

func (s *Sequence) RandomTime() time.Time {
	return time.Now().Add(time.Duration(s.RandomInt64()))
}
