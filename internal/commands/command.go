package commands

type Command interface {
	Run(params map[string]string) error
	GetName() string
	Help() string
}
