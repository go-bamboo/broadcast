package broadcast_test

import (
	"context"
	"testing"

	. "github.com/go-bamboo/broadcast"
)

func TestNewWatch(t *testing.T) {
	w := NewWatch[int](func(ctx context.Context, data int) error {
		t.Logf("%d", data)
		return nil
	}, 1)

	w.Start()
	for i := 0; i < 100; i++ {
		w.Send(i)
	}
	w.Stop()
}
