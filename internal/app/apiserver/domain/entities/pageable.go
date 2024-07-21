package entities

const (
	MinTotalItems = 0
	MinPageOrder  = 1
	MinLimitItems = 1
	MaxLimitItems = 25
)

type Pageable interface {
	SetPage(int) error
	SetLimit(int) error
	SetTotal(int) error
	GetPage() int
	GetLimit() int
	GetNext() int
	GetPrev() int
	GetTotal() int
	GetTotalPages() int
}

type EntityPageable struct {
	Page  int
	Limit int
	Total int
}

func NewEntityPageable(page int, limit int, total int) (*EntityPageable, error) {
	var err error
	entityPageable := EntityPageable{}

	err = entityPageable.SetPage(page)

	if err != nil {
		return nil, err
	}

	err = entityPageable.SetLimit(limit)

	if err != nil {
		return nil, err
	}

	err = entityPageable.SetTotal(total)

	if err != nil {
		return nil, err
	}

	return &entityPageable, nil
}

func (pageable *EntityPageable) SetPage(page int) error {
	pageable.Page = page

	if pageable.Page < MinPageOrder {
		pageable.Page = MinPageOrder
	}

	return nil
}

func (pageable *EntityPageable) SetLimit(limit int) error {
	pageable.Limit = limit

	if pageable.Limit < MinLimitItems {
		pageable.Limit = MaxLimitItems
	}

	if pageable.Limit > MaxLimitItems {
		pageable.Limit = MaxLimitItems
	}

	return nil
}

func (pageable *EntityPageable) SetTotal(total int) error {
	pageable.Total = total

	if pageable.Total < MinTotalItems {
		pageable.Total = MinTotalItems
	}

	return nil
}

func (pageable EntityPageable) GetPage() int {
	return pageable.Page
}

func (pageable EntityPageable) GetLimit() int {
	return pageable.Limit
}

func (pageable EntityPageable) GetTotal() int {
	return pageable.Total
}

func (pageable EntityPageable) GetNext() int {
	if pageable.GetPage() >= pageable.GetTotalPages() {
		return pageable.GetTotalPages()
	}

	return pageable.GetPage() + 1
}

func (pageable EntityPageable) GetPrev() int {
	if pageable.GetPage() > pageable.GetTotalPages() {
		return pageable.GetTotalPages()
	}

	if pageable.GetPage() > MinPageOrder {
		return pageable.GetPage() - 1
	}

	return pageable.GetPage()
}

func (pageable EntityPageable) GetTotalPages() int {
	totalPages := (pageable.Total + pageable.GetLimit() - 1) / pageable.GetLimit()

	if totalPages < MinPageOrder {
		return MinPageOrder
	}

	return totalPages
}
