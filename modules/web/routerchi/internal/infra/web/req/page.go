package req

import (
	"net/http"
	"strings"
)

// Based on:
// https://jsonapi.org/profiles/ethanresnick/cursor-pagination
// https://github.com/emicklei/go-restful/blob/master/request.go#L55
type PageRequest struct {
	Page   int `json:"page"`
	Size   int `json:"size"`
	Offset int `json:"offset"`
}

type SortRequest struct {
	Property  string `json:"property"`
	Direction string `json:"direction"`
}

func NewPage(r *http.Request) *PageRequest {
	page := DefaultQueryInt(r, "page", 1)
	if page < 1 {
		page = 1
	}

	size := DefaultQueryInt(r, "size", 20)
	if size < 1 {
		size = 20
	} else if size > 1000 {
		size = 1000
	}

	offset := size * (page - 1)
	return &PageRequest{
		Page:   page,
		Size:   size,
		Offset: offset,
	}
}

func NewSort(r *http.Request) []*SortRequest {
	values := QueryAll(r, "sort")
	if len(values) == 0 {
		return []*SortRequest{
			{
				Property:  "id",
				Direction: "asc",
			},
		}
	}

	var params []*SortRequest
	for _, val := range values {
		property, order, found := strings.Cut(val, ",")
		if found && (order == "asc" || order == "desc") {
			params = append(params, &SortRequest{
				Property:  property,
				Direction: order,
			})
		}
	}

	return params
}
