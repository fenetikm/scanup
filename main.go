package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Arguments:", os.Args[1:])

	// Check if arguments were provided
	if len(os.Args) > 1 {
		fmt.Println("First argument:", os.Args[1])
	}
}
