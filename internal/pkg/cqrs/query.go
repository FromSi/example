package cqrs

type Query interface {
}

type QueryHandler interface {
	Handle(Query) (any, error)
}

type QueryCQRS interface {
	Ask(Query) (any, error)
}
