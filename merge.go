package genbus

import "sync"

// Merge two or more channels with same type
func Merge[T any](input ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	out := make(chan T)

	output := func(c <-chan T) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(input))
	for _, c := range input {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
