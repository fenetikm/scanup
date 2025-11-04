package commands

import (
	"github.com/fenetikm/scanup/internal"
)

type Command interface {
	Run(state *internal.State) error
	Name() string
	Help() string
}
