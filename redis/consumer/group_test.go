package consumer

import (
	"github.com/shlima/oi/redis"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func empty(messageID string, values map[string]interface{}, ack AckFn) {
}

func TestGroup_Close(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		group := NewGroup(MustNewClearRedis(t))
		err := group.Register(empty, GroupOpts{
			Group:           "Group",
			Stream:          "Stream",
			Consumer:        "Consumer",
			OffsetPolicy:    OffsetNewest,
			BlockingTimeout: 0, // block forever
		})

		require.NoError(t, err)
		group.Close()
		err = group.Listen()
		require.NoError(t, err)
	})
}

func MustNewClearRedis(t *testing.T) *redis.Redis {
	rd, err := redis.Connect(os.Getenv("REDIS_DSN"))
	require.NoError(t, err)
	require.NoError(t, rd.FlushDB())
	return rd
}
