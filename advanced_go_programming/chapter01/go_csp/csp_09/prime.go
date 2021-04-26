package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := Generatenatural(ctx)
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v : %v\n", i+1, prime)
		PrimeFilter(ctx, ch, prime)
	}
	//time.Sleep(time.Second)
	cancel()

}

func Generatenatural(ctx context.Context) chan int {
	ch := make(chan int)

	go func() {
		for i := 2; ; i++ {
			select {
			case <-ctx.Done():
				return
			case ch <- i:
			}
		}
	}()
	return ch
}

func PrimeFilter(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)

	go func() {
		if i := <-in; i%prime != 0 {
			select {
			case <-ctx.Done():
				return
			case out <- i:
			}
		}
	}()
	return out
}
