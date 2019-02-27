package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)

	go func() {
		for {
			fmt.Print("-")
			time.Sleep(10 * time.Millisecond)
		}
	}()

	go func() {
		for {
			fmt.Print("~")
			time.Sleep(10 * time.Millisecond)
		}
	}()

	go func() {
		for {
			fmt.Print("*")
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	for {
		fmt.Print("_")
		sleep()
	}
}

func sleep() {
	for i := 0; i < 100000; i++ {
	}
	for i := 0; i < 100000; i++ {
	}
}
