package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

// goroutine, concept: it is a function executing concurrently with other goroutines in the same address space
// goroutine, keyword: go, then goroutine
// goroutine, communication: channel, see details channel-cheat-sheet.go
// goroutine, pass on: closed, so you need to pass in parameters
// goroutine, pass on: for variable, need to be passed in i to go fn
// goroutine, pass on: chan variable, need to be passed in chan to go fn
// goroutine, pass on: in case fn write ch, then parameters: chan<-
// goroutine, pass on: in case fn read ch, then parameters: <-chan
// goroutine, read an write global variables: need Mutex, m.Lock(), m.Unlock(), otherwise even if pointer atomicity
// goroutine, anonymous goroutine
// Processes, process, corresponding to multiple Threads, multiple Ports
// Thread, Thread

var ggg int

func main() {
	c := make(chan int)     // communicate
	go func(c chan<- int) { // formal parameters
		list := []int{3, 2, 1}
		sort.Ints(list)
		c <- 1
	}(c)
	time.Sleep(time.Second)
	go func() {
		list := []int{3, 2, 1}
		sort.Ints(list)
		c <- 2 // direct read and write chan
	}()

	m := sync.Mutex{} // mutex read and write global variables
	for i := 0; i < 999; i++ {
		go func() {
			m.Lock()
			ggg = ggg + 1
			m.Unlock()
		}()
	}

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(ggg)
}
