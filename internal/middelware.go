package internal

type Middelware interface {
	Exec(ctx *context, cmd Command) (next bool, err error)
}
