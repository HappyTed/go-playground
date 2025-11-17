package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("example.txt") // Go 1.16+
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
