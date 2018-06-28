package main

import (
	"fmt"
	"sync"
	"time"
)

// When main decides to exit without receiving all the values from out, it must tell the goroutines in the upstream stages to abandon the values they're trying to send. It does so by sending values on a channel called done. It sends two values since there are potentially two blocked senders:
func main() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the first value from output.
	done := make(chan struct{}, 2)
	out := merge(done, c1, c2)
	fmt.Println(<-out) // 4 or 9

	// Tell the remaining senders we're leaving.
	done <- struct{}{}
	done <- struct{}{}

}
func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()
	return out
}
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
			time.Sleep(time.Second * 5)
		}
		close(out)
	}()
	return out
}

// when the consume does not consume any more .we need to known
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {

	var wg sync.WaitGroup
	out := make(chan int)
	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed or it receives a value
	// from done, then output calls wg.Done.

	// This approach has a problem: each downstream receiver needs to know the number of potentially blocked upstream senders and arrange to signal those senders on early return. Keeping track of these counts is tedious and error-prone.
	output := func(c <-chan int) {
		for n := range c {
			select {
			case out <- n:
			case <-done:
			}
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
