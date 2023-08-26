package genbus

import "unsafe"

func Subscribe[T any](bus *PubSub) <-chan T {
	out := make(chan T)

	go func(c <-chan any) {
		for n := range c {
			out <- n.(T)
		}
		close(out)
	}(bus.Subscribe(new(T)))

	return out
}

func SubscribeAs[T, A any](bus *PubSub) <-chan T {
	out := make(chan T)

	go func(c <-chan any) {
		for n := range c {
			alias := n.(A)
			// shooting yourself in the foot, just cause
			desirable := (*T)(unsafe.Pointer(&alias))
			out <- *desirable
		}
		close(out)
	}(bus.Subscribe(new(A)))

	return out
}
