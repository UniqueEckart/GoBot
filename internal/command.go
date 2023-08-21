package internal

type Command interface {
	Invokes() []string
	Description() string
	HelpSyntax() string
	RequiredUserPermissions() int64
	AdminRequired() bool
	HasSubCommands() bool
	Exec(ctx Context) error
}
