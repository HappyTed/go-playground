package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func timeHamdler(w http.ResponseWriter, r *http.Request) {
	// TODO: заполнить функцию
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: заполнить функцию
}

func main() {
	PORT := ":8001"
	args := os.Args
	if len(args) != 1 {
		PORT = ":" + args[1]
	}
	fmt.Println("Using port number: ", PORT)

	mux := http.NewServeMux()
	serv := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		ReadTimeout:  3 * time.Second, // указывает максимальное время, разрешенное для чтения всего клиентского запроса, включая тело
		WriteTimeout: 3 * time.Second, // максимальная продолжительность времени до отправки ответа клиенту
	}

	mux.HandleFunc("/time", timeHamdler)
	mux.HandleFunc("/", myHandler)

	err := serv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
