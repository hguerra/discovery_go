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
	firstPage := 1
	defaultSize := 20
	minSize := 1
	maxSize := 1000

	page := DefaultQueryInt(r, "page", firstPage)
	if page < firstPage {
		page = firstPage
	}

	size := DefaultQueryInt(r, "size", defaultSize)
	if size < minSize {
		size = defaultSize
	} else if size > maxSize {
		size = maxSize
	}

	offset := size * (page - firstPage)
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
