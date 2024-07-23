package filters

type FindPostFilter struct {
	ID string
}

func NewFindPostFilter(id string) (*FindPostFilter, error) {
	return &FindPostFilter{
		ID: id,
	}, nil
}
