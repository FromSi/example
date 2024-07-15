package entities

import (
	"errors"
)

const (
	MinTotal      = 0
	MinPageOrder  = 1
	MinLimitItems = 1
	MaxLimitItems = 25
)

type Pageable interface {
	SetTotal(int)
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
	if total < MinTotal {
		return nil, errors.New("total value is below the minimum allowed value")
	}

	return &EntityPageable{
		Page:  page,
		Limit: limit,
	}, nil
}

func (pageable *EntityPageable) SetTotal(total int) {
	pageable.Total = total
}

func (pageable EntityPageable) GetPage() int {
	if pageable.Page < MinPageOrder {
		return MinPageOrder
	}

	return pageable.Page
}

func (pageable EntityPageable) GetLimit() int {
	if pageable.Limit < MinLimitItems {
		return MaxLimitItems
	}

	if pageable.Limit > MaxLimitItems {
		return MaxLimitItems
	}

	return pageable.Limit
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

func (pageable EntityPageable) GetTotal() int {
	return pageable.Total
}

func (pageable EntityPageable) GetTotalPages() int {
	totalPages := (pageable.Total + pageable.GetLimit() - 1) / pageable.GetLimit()

	if totalPages < MinPageOrder {
		return MinPageOrder
	}

	return totalPages
}
