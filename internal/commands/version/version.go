package commands

import (
	"fmt"
	"github.com/fenetikm/scanup/internal"
)

type Version struct{}

func (v *Version) Run(state *internal.State) error {
	fmt.Println("Version 0.0.1")
	return nil
}

func (v *Version) GetName() string {
	return "version"
}

func (v *Version) Help() string {
	// Todo: get the app name from config
	return "Returns the version of ScanUp"
}
