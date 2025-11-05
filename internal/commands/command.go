package commands

import (
	"flag"

	"github.com/fenetikm/scanup/internal"
)

type Command struct {
	Run         func(state *internal.State, args []string) (string, error)
	Name        string
	Description string
	flags       *flag.FlagSet
}

func CreateCommand(name string, description string) *Command {
	return &Command{
		Name:        name,
		Description: description,
	}
}
