package responses

import (
	"github.com/fromsi/example/internal/app/apiserver/domain/entities"
)

type Response struct {
	Data any `json:"data"`
}

type ListResponse struct {
	Data     any               `json:"data"`
	Pageable *PageableResponse `json:"pageable,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type PageableResponse struct {
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
	Next       int `json:"next"`
	Page       int `json:"page"`
	Prev       int `json:"prev"`
	Limit      int `json:"limit"`
}

func NewPageableResponse(pageable entities.Pageable) *PageableResponse {
	if pageable == nil {
		return nil
	}

	return &PageableResponse{
		Total:      pageable.GetTotal(),
		TotalPages: pageable.GetTotalPages(),
		Next:       pageable.GetNext(),
		Page:       pageable.GetPage(),
		Prev:       pageable.GetPrev(),
		Limit:      pageable.GetLimit(),
	}
}
