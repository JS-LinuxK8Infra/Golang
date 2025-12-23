package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("worker exiting")
				time.Sleep(100 * time.Millisecond)
				return
			}
		}
	}()

	time.Sleep(1 * time.Second)
	cancel()
	time.Sleep(100 * time.Millisecond)
}
