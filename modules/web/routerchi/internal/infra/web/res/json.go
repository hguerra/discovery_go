package res

import (
	"bytes"
	"encoding/json"
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

// JSON marshals 'v' to JSON, automatically escaping HTML and setting the
// Content-Type as application/json.
// Based on:
// https://github.com/go-chi/render/blob/master/responder.go#L93
// https://github.com/gmhafiz/go8/blob/master/internal/utility/respond/json.go
func JSON(w http.ResponseWriter, r *http.Request, status int, v any) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	/* #nosec G104 */
	w.Write(buf.Bytes()) //nolint:errcheck
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
