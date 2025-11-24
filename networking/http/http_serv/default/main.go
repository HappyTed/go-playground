package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The cyrrent time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func main() {
	PORT := ":8001"

	args := os.Args
	if len(args) != 1 {
		PORT = ":" + args[1]
	}

	fmt.Println("Using port:", PORT)

	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/", defaultHandler)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		panic(err)
	}

}
