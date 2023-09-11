package batcher

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

// ErrFull is an error indicating that the channel is full.
var ErrFull = errors.New("channel is full")

// Option is an interface for optional configurations.
type Option interface {
	apply(*options)
}

// options holds the configuration options for the Batcher.
type options struct {
	size     int           // Maximum batch size
	buffer   int           // Channel buffer size
	worker   int           // Number of worker goroutines
	interval time.Duration // Time interval for batching
}

// check sets default values for options if they are not set.
func (o options) check() {
	if o.size <= 0 {
		o.size = 100
	}
	if o.buffer <= 0 {
		o.buffer = 100
	}
	if o.worker <= 0 {
		o.worker = 5
	}
	if o.interval <= 0 {
		o.interval = time.Second
	}
}

// funcOption is an implementation of the Option interface.
type funcOption struct {
	f func(*options)
}

// apply applies the function to modify options.
func (fo *funcOption) apply(o *options) {
	fo.f(o)
}

// newOption creates a new funcOption.
func newOption(f func(*options)) *funcOption {
	return &funcOption{
		f: f,
	}
}

// WithSize sets the batch size.
func WithSize(s int) Option {
	return newOption(func(o *options) {
		o.size = s
	})
}

// WithBuffer sets the channel buffer size.
func WithBuffer(b int) Option {
	return newOption(func(o *options) {
		o.buffer = b
	})
}

// WithWorker sets the number of worker goroutines.
func WithWorker(w int) Option {
	return newOption(func(o *options) {
		o.worker = w
	})
}

// WithInterval sets the time interval for batching.
func WithInterval(i time.Duration) Option {
	return newOption(func(o *options) {
		o.interval = i
	})
}

// msg represents a message to be batched.
type msg struct {
	key string      // The key for sharding
	val interface{} // The value to be batched
}

// Batcher is the main struct that handles batching.
type Batcher struct {
	opts options // Configuration options

	Do       func(ctx context.Context, val map[string][]interface{}) // Function to process batches
	Sharding func(key string) int                                    // Function to determine sharding
	chans    []chan *msg                                             // Channels for each worker
	wait     sync.WaitGroup                                          // Wait group for goroutines
}

// New creates a new Batcher with optional configurations.
func New(opts ...Option) *Batcher {
	b := &Batcher{}
	for _, opt := range opts {
		opt.apply(&b.opts)
	}
	b.opts.check()

	// Initialize channels for each worker
	b.chans = make([]chan *msg, b.opts.worker)
	for i := 0; i < b.opts.worker; i++ {
		b.chans[i] = make(chan *msg, b.opts.buffer)
	}
	return b
}

// Start initializes the Batcher and starts worker goroutines.
func (b *Batcher) Start() {
	if b.Do == nil {
		log.Fatal("Batcher: Do func is nil")
	}
	if b.Sharding == nil {
		log.Fatal("Batcher: Sharding func is nil")
	}
	b.wait.Add(len(b.chans))
	for i, ch := range b.chans {
		go b.merge(i, ch)
	}
}

// Add adds a new message to the appropriate channel.
func (b *Batcher) Add(key string, val interface{}) error {
	ch, msg := b.add(key, val)
	select {
	case ch <- msg:
	default:
		return ErrFull
	}
	return nil
}

// add determines the appropriate channel and message for a given key and value.
func (b *Batcher) add(key string, val interface{}) (chan *msg, *msg) {
	sharding := b.Sharding(key) % b.opts.worker
	ch := b.chans[sharding]
	msg := &msg{key: key, val: val}
	return ch, msg
}

// merge is a worker function that merges messages into batches.
func (b *Batcher) merge(idx int, ch <-chan *msg) {
	defer b.wait.Done()

	// Initialize variables
	var (
		msg        *msg
		count      int
		closed     bool
		lastTicker = true
		interval   = b.opts.interval
		vals       = make(map[string][]interface{}, b.opts.size)
	)
	if idx > 0 {
		interval = time.Duration(int64(idx) * (int64(b.opts.interval) / int64(b.opts.worker)))
	}
	ticker := time.NewTicker(interval)
	for {
		select {
		case msg = <-ch:
			if msg == nil {
				closed = true
				break
			}
			count++
			vals[msg.key] = append(vals[msg.key], msg.val)
			if count >= b.opts.size {
				break
			}
			continue
		case <-ticker.C:
			if lastTicker {
				ticker.Stop()
				ticker = time.NewTicker(b.opts.interval)
				lastTicker = false
			}
		}
		if len(vals) > 0 {
			ctx := context.Background()
			b.Do(ctx, vals)
			vals = make(map[string][]interface{}, b.opts.size)
			count = 0
		}
		if closed {
			ticker.Stop()
			return
		}
	}
}

// Close stops all worker goroutines.
func (b *Batcher) Close() {
	for _, ch := range b.chans {
		ch <- nil
	}
	b.wait.Wait()
}
