package lesson20

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancell(ctx context.Context) bool {
	select {
	case <-ctx.Done(): // 接收取消通知
		return true
	default:
		return false
	}
}

func TestContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCancell(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancel")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second)
}
