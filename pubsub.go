package genbus

import (
	"reflect"
	"sync"
)

type PubSub struct {
	mu     sync.Mutex
	subs   map[reflect.Type][]chan any
	quit   chan struct{}
	closed bool
}

func New() *PubSub {
	return &PubSub{
		subs: make(map[reflect.Type][]chan any),
		quit: make(chan struct{}),
	}
}

// Publish is synchronous, it will lock until publishes in all subscriptions
func (b *PubSub) Publish(msg any) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return
	}

	// Kinda synchronous in some ways, beware
	for _, ch := range b.subs[reflect.TypeOf(msg)] {
		ch <- msg
	}
}

// Subscribe to typed topic
func (b *PubSub) Subscribe(to any) <-chan any {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return nil
	}

	ch := make(chan any)
	msgType := reflect.TypeOf(to).Elem()
	b.subs[msgType] = append(b.subs[msgType], ch)
	return ch
}

func (b *PubSub) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return
	}

	b.closed = true
	close(b.quit)

	for _, ch := range b.subs {
		for _, sub := range ch {
			close(sub)
		}
	}
}
