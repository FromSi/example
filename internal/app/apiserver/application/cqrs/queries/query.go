package queries

type Query interface {
}

type QueryHandler interface {
	Handle(Query) (any, error)
}
