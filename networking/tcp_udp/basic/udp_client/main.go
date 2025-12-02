package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// запуск сервера через nc: `nc -ul localhost 1234 -vv`
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a host:port string")
		return
	}
	CONNECT := arguments[1]

	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	if err != nil {
		log.Fatal(err)
	}

	c, err := net.DialUDP("udp4", nil, s) // инициализируем фактическое соединение
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())
	defer c.Close()

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')

		data := []byte(text + "\n")

		_, err := c.Write(data)
		if err != nil {
			log.Fatal("Fatal write data: ", err)
		}

		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting UDP client!")
			return
		}

		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			log.Fatal("fatal read from udp: ", err)
		}
		fmt.Printf("Reply: %s\n", string(buffer[0:n]))
	}
}
