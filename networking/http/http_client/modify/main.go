package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL\n", filepath.Base(os.Args[0]))
		return
	}

	URL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println("Error in parsing url:", err)
		return
	}

	// http client
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	// объект запроса
	request, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// выполняем запрос через http client
	httpData, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Request Status:", httpData.Status)
	fmt.Println("Status code:", httpData.StatusCode)

	// httputil.DumpResponse используется для получения ответа от сервера и в основном в целях отладки.
	// Второй аргумент говорит о том, включать ли в вывод
	// Для того что бы сделать тоже самое на стороне сервера, нужно использовать httputil.DumpRequest().
	header, _ := httputil.DumpResponse(httpData, false)
	fmt.Println(string(header))

	contentType := httpData.Header.Get("Content-Type")
	characherSet := strings.SplitAfter(contentType, "charset=")
	if len(characherSet) > 1 {
		fmt.Println("Character Set:", characherSet[1])
	}

	// длина содержимого ответа на запроса
	if httpData.ContentLength == -1 {
		fmt.Println("ContentLength is unknown!")
	} else {
		fmt.Println("ContentLength:", httpData.ContentLength)
	}

	length := 0
	var buffer [1024]byte
	r := httpData.Body
	for {
		n, err := r.Read(buffer[0:])
		if err != nil {
			fmt.Println(err)
			break
		}
		length = length + n
	}
	fmt.Println("Calculated response data length:", length)

	r.Close()
}
