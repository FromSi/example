package cqrs

type Command interface {
}

type CommandHandler interface {
	Handle(Command) error
}

type CommandCQRS interface {
	Dispatch(Command) error
}
