package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

/*
Для использования установи netcat: `sudo pacman -S netcat` или `sudo apt install netcat`
затем запусти сервер tcp: `nc -l localhost 1234 -vv`

В другом терминале запускаем клиент: `go run main.go localhost:1234`
для выхода из клиента отправляем в stdin 1) клиента `STOP` 2) в stdin сервера жмём enter (что бы произвести ответ от сервера)
*/
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a server:port string!")
		return
	}

	connect := arguments[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", connect)
	if err != nil {
		log.Fatal("ResolveTCPAddr:", err)
	}

	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		log.Fatal("DialTCP:", err) // TODO: падает в этом месте: "DialTCP:dial tcp 127.0.0.1:1234: connect: connection refused", нужно разбираться.
		// Не работает с netcat, но с ../tcp_serv/main.go работает
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			conn.Close()
			return
		}
	}
}
