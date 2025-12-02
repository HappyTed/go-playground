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
Для использования установи netcat: sudo pacman -S netcat || sudo apt install netcat
затем запусти сервер tcp: `nc -l localhost 1234 -vv`

запускаем клиент: `go run main.go localhost:1234`
для выхода из клиента отправляем в stdin 1) клиента `STOP` 2) в stdin сервера жмём enter (что бы произвести ответ от сервера)
*/
func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide host:port")
		return
	}

	connect := args[1]
	c, err := net.Dial("tcp", connect)
	if err != nil {
		log.Fatal(err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		// клиент ожидает ответа сервера, прежде чем завершить работу после ввода STOP
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
