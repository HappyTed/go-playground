package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	fmt.Println("Full path:", args[0])

	// получить последний элемент пути
	fmt.Println("Base path:", filepath.Base(args[0]))
}
