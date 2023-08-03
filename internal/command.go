package internal

type Command interface {
	Invokes() []string
	Descreption() string
	HelpSyntax() string
	AdminRequired() bool
	HasSubCommands() bool
	Exec(ctx Context) error
}
