package redis

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRedis_Streams(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		redis := MustNewClearRedis(t)
		message := map[string]interface{}{"bar": "baz"}

		err := redis.XAdd(&XAddArgs{Stream: "stream", Values: message})
		require.NoError(t, err)

		err = redis.XGroupCreateMkStreamSafe("stream", "group", "0")
		require.NoError(t, err)

		streams, err := redis.XReadGroup(&XReadGroupArgs{
			Consumer: "UNIQUE_CONSUMER_ID",
			Group:    "group",
			Count:    1,
			Streams:  []string{"stream", ">"},
		})

		require.NoError(t, err)
		require.Len(t, streams, 1, "streams")
		require.Len(t, streams[0].Messages, 1, "messages")
		require.Equal(t, message, streams[0].Messages[0].Values)

		err = redis.XAck("stream", "group", streams[0].Messages[0].ID)
		require.NoError(t, err)
	})
}

func TestRedis_XGroupCreateMkStreamSafe(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		redis := MustNewClearRedis(t)
		err1 := redis.XGroupCreateMkStreamSafe("stream", "group", "0")
		err2 := redis.XGroupCreateMkStreamSafe("stream", "group", "0")
		require.NoError(t, err1)
		require.NoError(t, err2)
	})
}

func TestRedis_Ping(t *testing.T) {
	t.Run("it works", func(*testing.T) {
		err := MustNewClearRedis(t).Ping()
		require.NoError(t, err)
	})

	t.Run("it errors", func(*testing.T) {
		err := ConnectOpts(&Options{Addr: "foo"}).Ping()
		require.Error(t, err)
	})
}

func TestRedis_Subscribe(t *testing.T) {
	t.Run("it works", func(t *testing.T) {
		redis := MustNewClearRedis(t)
		sub := redis.Subscribe("foo")

		err := redis.Publish("foo", nil)
		require.NoError(t, err)

		message := <-sub.Channel()
		require.Equal(t, "foo", message.Channel)
	})
}

func MustNewClearRedis(t *testing.T) *Redis {
	redis, err := Connect(os.Getenv("REDIS_DSN"))
	require.NoError(t, err)
	require.NoError(t, redis.FlushDB())
	return redis
}
