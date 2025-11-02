package commands

import "fmt"

type Version struct{}

func (v *Version) Run(params map[string]string) error {
	// TODO: get ver from somewhere easier
	fmt.Println("Version 0.0.1")
	return nil
}

func (v *Version) GetName() string {
	return "version"
}

func (v *Version) Help() string {
	return "Returns the version of ScanUp"
}
