package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func parseArgs(args []string) (start int, end int, f string, p string) {

	start, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("Start must be int! But received: %s of %T\n", args[2], args[2])
	}

	end, err = strconv.Atoi(args[3])
	if err != nil {
		fmt.Printf("Start must be int! But received: %s of %T\n", args[2], args[2])
	}

	f = args[1]
	p = args[4]

	return start, end, f, p
}

func createFile(f string) {
	fmt.Println(f)
}

func main() {

	args := os.Args

	if len(args) != 5 {
		fmt.Println("Usage: go run main.go [fileBaseName] [startInt] [stopInt] [directory]")
		return
	}

	var (
		fileName string
		start    int
		end      int
		path     string
	)
	start, end, fileName, path = parseArgs(args)

	fmt.Printf("from=%d to=%d file=%s dir=%s\n", start, end, fileName, path)

	fmt.Println()

	var wg sync.WaitGroup
	for i := start; i <= end; i++ {
		wg.Add(1)
		filepath := fmt.Sprintf("%s/%s%d", path, fileName, i)

		go func(f string) {
			defer wg.Done()
			createFile(f)
		}(filepath)
	}

	wg.Wait()
}
