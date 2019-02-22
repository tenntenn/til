package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	//var eg errgroup.Group
	eg, ctx := errgroup.WithContext(context.Background())

	eg.Go(func() error {
		return errors.New("error")
	})

	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			}
			time.Sleep(1 * time.Second)
		}
	})

	done := make(chan struct{})

	go func() {
		defer close(done)
		if err := eg.Wait(); err != nil {
			fmt.Println(err)
		}
	}()

	select {
	case <-done:
		fmt.Println("done")
	case <-time.After(5 * time.Second):
		fmt.Println("time out")
	}
}
