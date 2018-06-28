package main

import (
	"fmt"
	"time"
)

// is the channel is nunbuffed , the programm will deadlock !!!
func main() {
	c1 := make(chan int, 1)
	c1 <- 2
	for i := 0; i < 10; i++ {
		fmt.Println("hello word ")
		time.Sleep(time.Duration(time.Second))
	}
	go func() {
		fmt.Println(<-c1)
	}()
}
