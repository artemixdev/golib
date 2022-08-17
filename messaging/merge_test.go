package messaging

import (
	"context"
	"fmt"
	"sort"
	"testing"
	"time"
)

func TestMerge(t *testing.T) {
	src := make([]chan string, 3)
	dst := make(chan string, 3)
	srcCasted := make([]<-chan string, len(src))
	for i := range src {
		ch := make(chan string, 1)
		src[i] = ch
		srcCasted[i] = ch
	}

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	done := make(chan struct{})
	go func() {
		Merge[string](ctx, true, dst, srcCasted...)
		done <- struct{}{}
	}()

	go func() {
		for i, ch := range src {
			ch <- fmt.Sprintf("sample%d", i)
		}
	}()

	select {
	case <-time.After(110 * time.Millisecond):
		t.Fatalf("function stuck")
	case <-done:
	}

	values := make([]string, 0, len(src))
	for range src {
		values = append(values, <-dst)
	}

	sort.Strings(values)

	for i, value := range values {
		if value != fmt.Sprintf("sample%d", i) {
			t.Fatalf("wrong value: %s", value)
		}
	}
}
