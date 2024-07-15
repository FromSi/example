package commands

type Command interface {
}

type CommandHandler interface {
	Handle(Command) error
}
