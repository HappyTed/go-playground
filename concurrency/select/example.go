package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Object struct {
	inCh    chan int
	outCh   chan int
	timeout time.Duration
	wg      sync.WaitGroup
}

var ex *Object

func init() {
	ex = &Object{}
}

func main() {
	var ex *Object = &Object{}

	ex.wg.Add(1)

	go func() {
		toDoSmth(0, 2, ex.inCh, ex.outCh)
	}()

}

func toDoSmth(min, max int, inCh chan<- int, outCh <-chan int) {
	time.Sleep(time.Second)
	for {
		select {
		case inCh <- rand.Intn(max-min) + min:
		case v := <-outCh:
			fmt.Println("Ended! Val from channel is:", v)
			// return
		case <-time.After(4 * time.Second):
			fmt.Println("Time out!")
			return
		}
	}
}
