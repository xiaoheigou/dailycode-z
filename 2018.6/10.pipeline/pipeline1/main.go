package main

import (
	"fmt"
	"sync"
	"time"
)

//https://blog.golang.org/pipelines
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

// func main() {
// 	// Set up the pipeline.
// 	c := gen(2, 3)
// 	out := sq(c)

// 	// Consume the output.
// 	fmt.Println(<-out) // 4
// 	fmt.Println(<-out) // 9
// }

// func main() {
//     // Set up the pipeline and consume the output.
//     for n := range sq(sq(gen(2, 3))) {
//         fmt.Println(n) // 16 then 81
//     }
// }
func main() {
	in := gen(2, 3)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	c2 := sq(in)

	// Consume the merged output from c1 and c2.
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c is closed, then calls wg.Done.
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
