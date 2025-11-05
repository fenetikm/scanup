package version

import (
	"github.com/fenetikm/scanup/internal"
	"github.com/fenetikm/scanup/internal/commands"
)

func run(state *internal.State, args []string) (string, error) {
	return "Version " + state.Version, nil
}

func Create() *commands.Command {
	version := commands.CreateCommand(
		"version",
		"Display the version of the program",
	)

	version.Run = run

	return version
}
