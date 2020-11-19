package channel

import (
	"context"
	"fmt"
	"time"
)

// for loop, exit when i > 3
func Channel1() {
	queue := make(chan int, 3)
	for i := 0; i < 3; i++ {
		queue <- i
	}
	for i := 0; i < 3; i++ {
		fmt.Println(<-queue)
	}
}

// exit for range loop by close channel
func Channel2() {
	queue := make(chan int, 3)
	for i := 0; i < 3; i++ {
		queue <- i
	}
	// need to close channel befor using range channel
	close(queue)
	for ele := range queue {
		fmt.Println(ele)
	}
}

// exit for while loop by ctx.Done
func Channel3() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(ctx context.Context) {
		for {
			select {
			case <-time.After(time.Second * 100):
				// exit the select loop
			case <-ctx.Done():
				fmt.Println("normal job run return, so exit for loop")
				return
			}
		}

	}(ctx)
}

// receive only channel should'nt close, direct compilation error. but send only channel can close normaly
func Channel4() {
	sendOnly := make(chan<- int, 3)
	for i := 0; i < 3; i++ {
		sendOnly <- i
	}
	close(sendOnly)

	receiveOnly := make(<-chan int, 3)
	for i := 0; i < 3; i++ {
		fmt.Println(<-receiveOnly)
	}
	// close(receiveOnly) error
}

// send to a close channel, then panic
func Channel5() {
	ch := make(chan int, 0)
	close(ch)
	ch <- 1
}

// receive a close channel, then default value
func Channel6() {
	ch := make(chan int, 0)
	close(ch)
	<-ch
}
