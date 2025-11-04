package internal

import (
	"github.com/fenetikm/scanup/internal/config"
)

type State struct {
	Name    string
	Version string
	Config  *config.Config
}
