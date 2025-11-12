package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var myUrl string
var delay int = 5
var wg sync.WaitGroup

type myData struct {
	r   *http.Response
	err error
}

func connect(c context.Context) error {
	defer wg.Done()
	data := make(chan myData, 1)
	tr := &http.Transport{}
	httpClient := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", myUrl, nil)

	go func() {
		response, err := httpClient.Do(req)
		if err != nil {
			fmt.Println(err)
			data <- myData{nil, err}
		} else {
			pack := myData{response, err}
			data <- pack
		}
	}()

	select {
	case <-c.Done(): // подразумевается timeout контекста, что бы отменить клиентское соединение
		tr.CancelRequest(req) // TODO: адаптировать код под современный стиль 1.25+
		<-data
		fmt.Println("The request was cancaled!")
		return c.Err()
	case ok := <-data:
		err := ok.err
		resp := ok.r
		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}
		defer resp.Body.Close()

		raelHTTPData, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}
		fmt.Printf("Server Response: %s\n", raelHTTPData)
	}
	return nil
}

// examples:
// go run main.go http://localhost:80 5 # not ok
// go run main.go http://ya.ru:80 1 # will canceled
// go run main.go http://ya.ru:80 # 5 is default value. This OK
func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need a URL and a delay!")
		return
	}
	myUrl = os.Args[1]
	if len(os.Args) == 3 {
		t, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		delay = t
	}

	fmt.Println("Delay:", delay)
	c := context.Background()
	c, cancel := context.WithTimeout(c, time.Duration(delay)*time.Second)
	defer cancel()

	fmt.Printf("Connecting to %s \n", myUrl)
	wg.Add(1)
	go connect(c)
	wg.Wait()
	fmt.Println("Exiting...")
}
