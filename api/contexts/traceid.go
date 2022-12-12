package contexts

import (
	"context"
	"sync"
)

var (
	logNo int = 1
	mu    sync.Mutex
)

type traceIDKey struct{}

func NewTraceID() int {
	var no int

	mu.Lock()
	no = logNo
	logNo += 1
	mu.Unlock()

	return no
}

func GetTracdID(ctx context.Context) int {
	v := ctx.Value(traceIDKey{})

	if id, ok := v.(int); ok {
		return id
	}

	return 0
}

func SetTraceID(ctx context.Context, tracdID int) context.Context {
	return context.WithValue(ctx, traceIDKey{}, tracdID)
}
