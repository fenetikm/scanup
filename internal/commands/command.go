package commands

type Command interface {
	run(params map[string]string) error
	getName() string
}
