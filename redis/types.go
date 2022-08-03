package redis

import "github.com/go-redis/redis"

type (
	Client          = redis.Client
	Options         = redis.Options
	XAddArgs        = redis.XAddArgs
	XReadGroupArgs  = redis.XReadGroupArgs
	XStreamSliceCmd = redis.XStreamSliceCmd
	XStream         = redis.XStream
	PubSub          = redis.PubSub
)

const (
	Nil = redis.Nil
)

type IRedis interface {
	XAdd(*XAddArgs) error
	FlushDB() error
	Ping() error

	// XGroupCreateMkStream https://redis.io/commands/xgroup-create/
	// This command creates a new consumer group uniquely identified by <consumersGroup>
	// for the stream stored at <subject>.
	// Every group has a unique name in a given stream.
	// When a consumer group with the same name already exists, the command returns a -BUSYGROUP error.
	// XGROUP CREATE stream group 0
	// XGROUP CREATE stream group $ MKSTREAM
	// MKSTREAM is optional. It creates the stream if it doesn’t already exist)
	XGroupCreateMkStream(subject, consumersGroup, start string) error

	// XGroupCreateMkStreamSafe not raises an error if group hase been already created
	XGroupCreateMkStreamSafe(stream, consumersGroup, start string) error

	// XReadGroup
	// []string{"stream", ">"} will read all messages from the beginning of years
	// []string{"stream", "$"} will read only new messages from the first call of READ
	// (on first consumer connection)
	//
	// it is possible to create groups of clients that consume different parts of the messages
	// arriving in a given stream. If, for instance, the stream gets the new entries A, B,
	// and C and there are two consumers reading via a consumer group, one client will get,
	// for instance, the messages A and C, and the other the message B, and so forth.
	//
	// The NOACK subcommand can be used to avoid adding the message to the PEL in cases
	// where reliability is not a requirement and the occasional message loss is acceptable.
	// This is equivalent to acknowledging the message when it is read.
	XReadGroup(args *XReadGroupArgs) ([]XStream, error)

	XAck(stream, group string, ids ...string) error

	Publish(channel string, message interface{}) error
	Subscribe(channels ...string) *PubSub
}
