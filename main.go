package main

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/fenetikm/scanup/internal"
	"github.com/fenetikm/scanup/internal/commands/version"
	"github.com/fenetikm/scanup/internal/config"
)

//go:embed semver
var appVersion string
var appName string = "ScanUp!"

func main() {
	config, err := config.Read()
	if err != nil {
		panic(err)
	}
	var appState = internal.State{
		Name:    appName,
		Version: appVersion,
		Config:  &config,
	}

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatalf("Missing command argument.")
	}

	command := args[0]
	extraArgs := []string{}
	if len(args) > 1 {
		extraArgs = args[1:]
	}

	// Todo: replace with handlers/registry
	if command == "version" {
		cmd := version.Create()
		output, err := cmd.Run(&appState, extraArgs)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(output)
	}
}
