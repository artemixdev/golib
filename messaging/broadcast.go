package messaging

import (
	"context"
	"sync"
)

// Broadcast reads the source and forwards the value to all the destinations.
// Durable flag provides the delivery if canceled.
func Broadcast[T any](ctx context.Context, durable bool, src <-chan T, dst ...chan<- T) {
	for {
		select {
		case <-ctx.Done():
			return

		case value, ok := <-src:
			if !ok {
				return
			}

			wg := sync.WaitGroup{}

			for _, ch := range dst {
				wg.Add(1)

				go func(ch chan<- T) {
					defer wg.Done()

					if durable {
						ch <- value
						return
					}

					select {
					case <-ctx.Done():
					case ch <- value:
					}
				}(ch)
			}

			wg.Wait()
		}
	}
}
