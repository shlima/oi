package redis

import (
	"fmt"
	"strings"

	"github.com/go-redis/redis"
)

type Redis struct {
	Cli *Client
}

func New(cli *Client) *Redis {
	return &Redis{Cli: cli}
}

func Connect(dsn string) (*Redis, error) {
	opts, err := redis.ParseURL(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis dsn: %w", err)
	}

	return ConnectOpts(opts), nil
}

func ConnectOpts(opts *Options) *Redis {
	return New(redis.NewClient(opts))
}

func (r *Redis) XAdd(args *XAddArgs) error {
	return r.Cli.XAdd(args).Err()
}

func (r *Redis) FlushDB() error {
	return r.Cli.FlushDB().Err()
}

func (r *Redis) Ping() error {
	return r.Cli.Ping().Err()
}

func (r *Redis) XGroupCreateMkStream(stream, consumersGroup, start string) error {
	return r.Cli.XGroupCreateMkStream(stream, consumersGroup, start).Err()
}

func (r *Redis) XGroupCreateMkStreamSafe(stream, consumersGroup, start string) error {
	err := r.XGroupCreateMkStream(stream, consumersGroup, start)
	if err != nil && strings.HasPrefix(err.Error(), "BUSYGROUP") {
		return nil
	}

	return err
}

func (r *Redis) XReadGroup(args *XReadGroupArgs) ([]XStream, error) {
	return r.Cli.XReadGroup(args).Result()
}

func (r *Redis) XAck(stream, group string, ids ...string) error {
	return r.Cli.XAck(stream, group, ids...).Err()
}

func (r *Redis) Publish(channel string, message interface{}) error {
	return r.Cli.Publish(channel, message).Err()
}

func (r *Redis) Subscribe(channels ...string) *PubSub {
	return r.Cli.Subscribe(channels...)
}
