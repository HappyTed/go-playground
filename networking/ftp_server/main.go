package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT = ":8765"

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	w.WriteHeader(http.StatusOK)
	Body := "Thanks for visiting!\n"
	fmt.Fprintf(w, "%s", Body)
}

// Делаем доступыми файлы в директории /tmp/
// Для тестирования, открыть http://localhost:8765/static/
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", defaultHandler)

	fileServer := http.FileServer(http.Dir("/tmp"))
	// регестрируем файловы сервер на для всех url, начинающихся со static
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	fmt.Println("Starting server on:", PORT)
	err := http.ListenAndServe(PORT, mux)
	fmt.Println(err)
}
