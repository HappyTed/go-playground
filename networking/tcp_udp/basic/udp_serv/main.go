package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}
	PORT := ":" + arguments[1]

	addr, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	serv, err := net.ListenUDP("udp4", addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer serv.Close()
	fmt.Printf("Start listen on %s!\n", addr.String())

	buffer := make([]byte, 1024)

	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := serv.ReadFromUDP(buffer)
		fmt.Print("-> ", string(buffer[0:n-1]))

		for strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}

		data := []byte(strconv.Itoa(random(1, 1001)))
		fmt.Printf("$data: %s\n", data)

		_, err = serv.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
