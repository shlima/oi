package pipe

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type Pipelined struct {
	*Pipe[Pipelined]
	Values []int
}

func (p *Pipelined) Call(val int, err error) *Pipelined {
	return p.Piping(func() error {
		if err != nil {
			return err
		}

		p.Values = append(p.Values, val)
		return nil
	})
}

func TestPipe(t *testing.T) {
	t.Run("when no errors", func(t *testing.T) {
		p := MustNewPipelined(t)
		p.Call(1, nil).Call(2, nil).Call(3, nil)
		require.Equal(t, []int{1, 2, 3}, p.Values)
		require.NoError(t, p.Error())
	})

	t.Run("when error occur", func(t *testing.T) {
		p := MustNewPipelined(t)
		err := errors.New("foo")
		p.Call(1, nil).Call(2, err).Call(3, nil)
		require.Equal(t, []int{1}, p.Values)
		require.ErrorIs(t, p.Error(), err)
	})
}

func MustNewPipelined(t *testing.T) *Pipelined {
	out := new(Pipelined)
	out.Pipe = New[Pipelined](out)
	return out
}
