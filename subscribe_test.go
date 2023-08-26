package genbus

import (
	"testing"
)

type source struct {
	Something string
}

func TestSubscribe(t *testing.T) {
	bus := New()

	sub := Subscribe[source](bus)
	go bus.Publish(source{Something: "test"})

	if res := <-sub; res.Something != "test" {
		t.Errorf("bus returned %s, want %s", res.Something, "test")
	}
}

type aliasedSource source

func TestSubscribeAs(t *testing.T) {
	bus := New()

	sub := SubscribeAs[source, aliasedSource](bus)
	go bus.Publish(aliasedSource{Something: "test"})

	// Getting `source` instead of `aliasedSource`
	if res := <-sub; res.Something != "test" {
		t.Errorf("bus returned %s, want %s", res.Something, "test")
	}
}
