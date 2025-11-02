package main

import (
	"log"
	"os"

	"github.com/fenetikm/scanup/internal/commands/version"
)

func main() {
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
