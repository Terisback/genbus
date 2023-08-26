package genbus

import (
	"fmt"
	"testing"
)

type Event struct {
	ID string
}

func TestPubSub(t *testing.T) {
	bus := New()

	// Subscribe to topic via type instance (types are used as topic)
	sub := bus.Subscribe(new(Event))

	// Send stuff
	go bus.Publish(Event{ID: "ready"})

	// Get stuff
	result := <-sub
	fmt.Println(result.(Event).ID) // `ready`
}
