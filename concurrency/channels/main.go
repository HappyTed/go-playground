package main

import (
	"fmt"
	"sync"
)

func writeIntToChannel(c chan int, x int) {
	c <- x
	close(c)
}

func writeTrueToChannel(c chan bool) {
	c <- true
}

func main() {
	// буферизированный канал
	c := make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(1)
	go func(c chan int) {
		defer wg.Done()
		writeIntToChannel(c, 10)
		fmt.Println("exit from goroutine after write to buffer channel")
	}(c)

	fmt.Println("Read from channel:", <-c)

	_, ok := <-c
	if ok {
		fmt.Println("Channel is open")
	} else {
		fmt.Println("Channel is closed")
	}

	wg.Wait()

	// не буферизированный канал
	var cb chan bool = make(chan bool)
	for i := 0; i < 5; i++ {
		go writeTrueToChannel(cb)
	}

	// диапазон по каналам
	// ВАЖНО: поскольку канал cb не закрыт‚
	// цикл по диапазону не завершается сам по себе
	n := 0
	for b := range cb {
		fmt.Println(b)
		if b == true {
			n++
		}
		if n > 2 {
			fmt.Println("n:", n)
			close(cb)
			break
		}
	}

	// чтение из закрытого канала
	for i := 0; i < 5; i++ {
		fmt.Println(<-cb)
	}

}
