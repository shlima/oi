package consumer

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/shlima/oi/redis"
)

// GroupOpts refs https://github.com/robinjoseph08/redisqueue/blob/master/consumer.go
type GroupOpts struct {
	Stream       string
	Group        string
	Consumer     string
	OffsetPolicy Offset

	// BlockingTimeout designates how long the XREADGROUP call blocks for. If
	// this is 0, it will block indefinitely. While this is the most efficient
	// from a polling perspective, if this call never times out, there is no
	// opportunity to yield back to Go at a regular interval. This means it's
	// possible that if no messages are coming in, the consumer cannot
	// gracefully shutdown. Instead, it's recommended to set this to 1-5
	// seconds, or even longer, depending on how long your application can wait
	// to shutdown.
	BlockingTimeout time.Duration
}

type Group struct {
	redis   redis.IRedis
	closed  bool
	mx      *sync.Mutex
	wg      *sync.WaitGroup
	handler HandlerFn
	args    *redis.XReadGroupArgs
}

func NewGroup(redis redis.IRedis) *Group {
	return &Group{
		redis: redis,
		mx:    new(sync.Mutex),
		wg:    new(sync.WaitGroup),
	}
}

// Close used for gracefully stop consumer
func (g *Group) Close() {
	g.mx.Lock()
	g.closed = true
	g.mx.Unlock()
	g.wg.Wait()
}

func (g *Group) Register(handler HandlerFn, opts GroupOpts) error {
	g.mx.Lock()
	defer g.mx.Unlock()

	if err := g.redis.XGroupCreateMkStreamSafe(opts.Stream, opts.Group, "0"); err != nil {
		return fmt.Errorf("failed to create consumer group: %w", err)
	}

	args := &redis.XReadGroupArgs{
		Count:    1,
		Group:    opts.Group,
		Consumer: opts.Consumer,
		Streams:  []string{opts.Stream, string(opts.OffsetPolicy)},
		Block:    opts.BlockingTimeout,
	}

	g.args = args
	g.handler = handler

	return nil
}

func (g *Group) IsClosed() bool {
	g.mx.Lock()
	defer g.mx.Unlock()
	return g.closed
}

func (g *Group) Listen() error {
	if g.handler == nil {
		return fmt.Errorf("please register consumer handler")
	}

	if g.args == nil {
		return fmt.Errorf("please register consumer args")
	}

	return g.listen()
}

func (g *Group) listen() error {
	g.wg.Add(1)
	defer g.wg.Done()

	for {
		if g.IsClosed() {
			return nil
		}

		streams, err := g.redis.XReadGroup(g.args)
		if err != nil {
			if isErrContinue(err) {
				continue
			}

			return err
		}

		for si := range streams {
			stream := streams[si]
			for mi := range stream.Messages {
				message := stream.Messages[mi]
				g.handler(message.ID, message.Values, func() error {
					return g.redis.XAck(stream.Stream, g.args.Group, message.ID)
				})
			}
		}
	}
}

func isErrContinue(err error) bool {
	if e, ok := err.(net.Error); ok && e.Timeout() {
		return true
	}

	if err == redis.Nil {
		return true
	}

	return false
}
