package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	SERVER := "localhost:" + arguments[1]
	s, err := net.ResolveTCPAddr("tcp", SERVER)
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.ListenTCP("tcp", s)
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 1024)
	conn, err := l.Accept()
	if err != nil {
		log.Fatal(err)
	}

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Fatal(err)
		}

		if strings.TrimSpace(string(buffer[:n])) == "STOP" {
			fmt.Println("Exiting TCP server !")
			conn.Close()
			return
		}

		fmt.Print("> ", string(buffer[:n-1]), "\n")
		_, err = conn.Write(buffer)
		if err != nil {
			log.Fatal(err)
		}
	}
}
