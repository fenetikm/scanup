package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fenetikm/scanup/internal/commands/version"
	"github.com/fenetikm/scanup/internal/config"
)

func main() {
	config, err := config.Read()
	if err != nil {
		panic(err)
	}
	fmt.Println(config)

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalf("Missing command argument.")
	}

	if args[0] == "version" {
		cmd := commands.Version{}
		err := cmd.Run(nil)
		if err != nil {
			log.Fatal(err)
		}
	}
}
