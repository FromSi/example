package filters

type FindUserFilter struct {
	ID string
}

func NewFindUserFilter(id string) (*FindUserFilter, error) {
	return &FindUserFilter{
		ID: id,
	}, nil
}
