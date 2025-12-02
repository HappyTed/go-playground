package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var count = 0

func handleConnection(c net.Conn) {
	me := count
	fmt.Printf("\\(>ω<)/\tlisten by: %d\n", me) // информарование о новом подключении

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		text := strings.TrimSpace(netData)
		if text == "STOP" {
			fmt.Printf("STOPING %d...\n", me)
			break
		}
		fmt.Printf("$%d: %s\n", me, text)

		answer := "Client number: " + strconv.Itoa(me) + "\n"
		c.Write([]byte(answer))
	}
	c.Close()
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}
	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	fmt.Printf("Start listen on %s!\n", PORT)

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(c)
		count++
	}
}
