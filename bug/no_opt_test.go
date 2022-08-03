package bug

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewNoOpt(t *testing.T) {
	t.Run("it implements IBug", func(t *testing.T) {
		require.Implements(t, (*IBug)(nil), new(NoOpt))
	})
}

func TestNoOpt_Ensure(t *testing.T) {
	t.Run("it does nothing", func(t *testing.T) {
		NewNoOpt().Ensure(errors.New("foo"))
		NewNoOpt().Ensure(nil)
	})
}
