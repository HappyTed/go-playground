package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Введите url")
		return
	}

	url := os.Args[1]

	data, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(os.Stdout, data.Body) // считывает данные  из data.Body и записывает их в io.Writer
	if err != nil {
		log.Fatal(err)
	}
	defer data.Body.Close() // закрываем data.Body, что бы облегчить работу сборщику мусора
}
