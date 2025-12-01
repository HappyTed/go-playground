package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]

	l, err := net.Listen("tcp", PORT) // Если второй параметр функции содержит номер порта без IP-адреса или имени хоста,
	// то она прослушивает все доступные IP-адреса локальной системы, как в данном случае
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	conn, err := l.Accept() // В этом конкретном TCP-сервереесть кое-что необычное:
	// он может обслуживать только первого TCP-клиента, который собирается подключиться к нему,
	// поскольку вызов Accept() находится вне цикла for и поэтому вызывается лишь единожды.
	// Каждый отдельный клиент должен быть указан с помощью другого вызова Accept()
	if err != nil {
		log.Fatal(err)
	}

	for {
		netData, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if strings.TrimSpace(netData) == "STOP" {
			fmt.Println("Exiting TCP server!")
			return
		}

		fmt.Print("-> ", netData)
		t := time.Now()
		myTime := t.Format(time.RFC3339) + "\n"
		conn.Write([]byte(myTime))
	}
}
