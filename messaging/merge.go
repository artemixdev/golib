package messaging

import (
	"context"
	"sync"
)

// Merge reads all the sources and forwards the values to the destination.
// Durable flag provides the delivery if canceled.
func Merge[T any](ctx context.Context, durable bool, dst chan<- T, src ...<-chan T) {
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for _, ch := range src {
		wg.Add(1)

		go func(ch <-chan T) {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return

				case value, ok := <-ch:
					if !ok {
						return
					}

					if durable {
						dst <- value
						continue
					}

					select {
					case <-ctx.Done():
						return
					case dst <- value:
					}
				}
			}
		}(ch)
	}
}
