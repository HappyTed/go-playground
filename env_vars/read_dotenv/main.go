package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(".env")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "=")
		if len(items) != 2 {
			panic("error reading .env: error in write format")
		}
		os.Setenv(items[0], strings.Replace(items[1], "\"", "", -1))
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	}
}
