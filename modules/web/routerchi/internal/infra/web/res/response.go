package res

import (
	"net/http"
)

// M is a convenience alias for quickly building a map structure that is going
// out to a responder. Just a short-hand.
type M map[string]any

// Based on:
// https://github.com/microsoft/api-guidelines/blob/vNext/Guidelines.md
// https://jsonapi.org/format/#document-top-level
// https://jsonapi.org/profiles/ethanresnick/cursor-pagination
// https://github.com/interagent/http-api-design/tree/master/en
// https://geemus.gitbooks.io/http-api-design/content/en/responses/return-appropriate-status-codes.html
type Page struct {
	Page  int `json:"page,omitempty"`
	Size  int `json:"size,omitempty"`
	Total int `json:"total,omitempty"`
}

type Meta struct {
	Page Page `json:"page,omitempty"`
}

type Response struct {
	Meta any `json:"meta,omitempty"`
	Data any `json:"data"`
}

type errDTO struct {
	Status  int    `json:"status,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
}

type ErrResponse struct {
	Error   errDTO   `json:"error"`
	Details []errDTO `json:"details,omitempty"`
}

func NewResponse(data any) Response {
	return Response{
		Data: data,
	}
}

func NewPageableResponse(data any, page, size, total int) Response {
	return Response{
		Meta: Meta{
			Page: Page{
				Page:  page,
				Size:  size,
				Total: total,
			},
		},
		Data: data,
	}
}

func NewErrResponse(status int, code, message string, errs ...string) *ErrResponse {
	if len(errs) == 0 {
		return &ErrResponse{
			Error: errDTO{
				Status:  status,
				Code:    code,
				Message: message,
			},
		}
	}

	var details []errDTO
	for _, e := range errs {
		details = append(details, errDTO{
			Message: e,
		})
	}

	return &ErrResponse{
		Error: errDTO{
			Status:  status,
			Code:    code,
			Message: message,
		},
		Details: details,
	}
}

func NewErr(w http.ResponseWriter, err *ErrResponse) {
	JSON(w, err.Error.Status, err)
}
