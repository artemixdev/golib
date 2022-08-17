package messaging

import (
	"context"
	"testing"
	"time"
)

func TestBroadcast(t *testing.T) {
	src := make(chan string)
	dst := make([]chan string, 3)
	dstCasted := make([]chan<- string, len(dst))
	for i := range dst {
		ch := make(chan string, 1)
		dst[i] = ch
		dstCasted[i] = ch
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	done := make(chan struct{})
	go func() {
		Broadcast[string](ctx, true, src, dstCasted...)
		done <- struct{}{}
	}()

	go func() {
		src <- "sample"
	}()

	select {
	case <-time.After(110 * time.Millisecond):
		t.Fatalf("function stuck")
	case <-done:
	}

	for _, ch := range dst {
		value := <-ch
		if value != "sample" {
			t.Fatalf("invalid forwarded value: %s", value)
		}
	}
}
