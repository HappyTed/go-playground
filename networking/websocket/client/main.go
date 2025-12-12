package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

var (
	SERVER       = ""
	PATH         = ""
	TIMESWAIT    = 0
	TIMESWAITMAX = 5
	in           = bufio.NewReader(os.Stdin)
)

// чтение stdin, запись в канал
func getInput(input chan string) {
	result, err := in.ReadString('\n') // блокирующая операция
	if err != nil {
		log.Println(err)
		return
	}
	input <- result
}

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("Need SERVER + PATH!")
		return
	}
	SERVER = arguments[1]
	PATH = arguments[2]
	fmt.Println("Connecting to:", SERVER, "at", PATH)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	input := make(chan string, 1)
	go getInput(input) // сразу же читаем Stdin

	URL := url.URL{Scheme: "ws", Host: SERVER, Path: PATH}
	c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil) // с помощью этого вызова создаётся подключение к ws
	if err != nil {
		log.Println("Error:", err)
		return
	}

	defer c.Close()

	done := make(chan struct{}) // сигнальный канал

	// чтение ответа от сервера
	go func() {
		defer close(done)

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("ReadMessage() error:", err)
				return
			}
			log.Printf("Received: %s", message)
		}
	}()

	// основой блок
	for {
		select {
		case <-time.After(4 * time.Second):
			log.Println("Please give me input!", TIMESWAITMAX-TIMESWAIT)
			TIMESWAIT++
			if TIMESWAIT > TIMESWAITMAX {
				syscall.Kill(syscall.Getpid(), syscall.SIGINT)
			}
		case <-done:
			return
		case t := <-input:
			err := c.WriteMessage(websocket.TextMessage, []byte(t))
			if err != nil {
				log.Println("Write error:", err)
				return
			}
			TIMESWAIT = 0
			go getInput(input) // блокирующая операция
		case <-interrupt:
			log.Println("Caught interrupt signal - quitting!")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Write close error:", err)
				return
			}
			// вот это непонятно
			select {
			case <-done:
			case <-time.After(2 * time.Second):
			}
			return
		}
	}
}
