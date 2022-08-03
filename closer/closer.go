package closer

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

type Closer struct {
	watching bool
	ch       chan os.Signal
	mx       *sync.Mutex
	signals  []os.Signal
	closers  []CloseFnE
}

func New(signals ...os.Signal) *Closer {
	return &Closer{
		mx:      new(sync.Mutex),
		signals: signals,
		// non-blocking to be able to detect second signal and quit immediately
		ch:       make(chan os.Signal, 1),
		closers:  make([]CloseFnE, 0),
		watching: false,
	}
}

func (c *Closer) WatchASync() {
	go c.WatchSync()
}

func (c *Closer) WatchSync() {
	c.watch(true)
	defer c.watch(false)
	defer c.mx.Unlock()

	defer signal.Stop(c.ch)
	signal.Notify(c.ch, c.signals...)

	for sig := range c.ch {
		go func(s os.Signal) {
			// it means another signal comes, when we are waiting for close
			if !c.mx.TryLock() {
				fmt.Printf("[⛔] double catch %s os signal, terminating\n", s)
				os.Exit(1)
			}

			fmt.Printf("[🌔] catch %s os signal, starting gracefull shutdown\n", s)

			switch c.close() {
			case true:
				os.Exit(0)
			default:
				os.Exit(1)
			}
		}(sig)
	}
}

func (c *Closer) close() (success bool) {
	success = true
	for ix := range c.closers {
		if err := c.closers[ix](); err != nil {
			success = false
			fmt.Printf("[⚠️] closer finished with error: %s\n", err.Error())
		}
	}

	return
}

func (c *Closer) AddE(fn CloseFnE) {
	c.mx.Lock()
	c.closers = append(c.closers, fn)
	c.mx.Unlock()
}

func (c *Closer) Add(fn CloseFn) {
	c.AddE(func() error {
		fn()
		return nil
	})
}

func (c *Closer) watch(value bool) {
	c.mx.Lock()
	c.watching = value
	c.mx.Unlock()
}

func (c *Closer) AutoWatchAsync() {
	c.mx.Lock()
	defer c.mx.Unlock()

	if c.watching {
		return
	}

	c.WatchASync()
}
