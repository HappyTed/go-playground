package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var timeout = time.Duration(time.Second * 2)

func Timeout(network, host string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, host, timeout)
	if err != nil {
		return nil, err
	}

	conn.SetDeadline(time.Now().Add(timeout))
	return conn, nil
}

func main() {

	t := http.Transport{
		Dial: Timeout,
	}
	client := http.Client{
		Transport: &t,
	}

	fmt.Println("Timeout value:", timeout)
	// TODO: smth with client
	// успех

	resp, err := client.Get("https://ya.ru") // успех
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	fmt.Printf("\n\t---\n")

	resp, err = client.Get("http://localhost:80") // выход по таймауту
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

}
