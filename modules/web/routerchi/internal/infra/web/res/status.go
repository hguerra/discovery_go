package res

import (
	"net/http"
)

func OK(w http.ResponseWriter, r *http.Request, v any) {
	JSON(w, r, http.StatusOK, NewResponse(v))
}

func PageOf(w http.ResponseWriter, r *http.Request, v any, page, size, total int) {
	JSON(w, r, http.StatusOK, NewPageableResponse(v, page, size, total))
}

func Created(w http.ResponseWriter, r *http.Request, v any) {
	JSON(w, r, http.StatusCreated, NewResponse(v))
}

func Accepted(w http.ResponseWriter, r *http.Request, v any) {
	JSON(w, r, http.StatusAccepted, NewResponse(v))
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func PartialContent(w http.ResponseWriter, r *http.Request, v any) {
	JSON(w, r, http.StatusPartialContent, NewResponse(v))
}

func BadRequest(w http.ResponseWriter, r *http.Request, code, message string) {
	NewErr(w, r, NewErrResponse(http.StatusBadRequest, code, message))
}

func Unauthorized(w http.ResponseWriter, r *http.Request, code, message string) {
	NewErr(w, r, NewErrResponse(http.StatusUnauthorized, code, message))
}

func Forbidden(w http.ResponseWriter, r *http.Request, code, message string) {
	NewErr(w, r, NewErrResponse(http.StatusForbidden, code, message))
}

func NotFound(w http.ResponseWriter, r *http.Request, code, message string) {
	NewErr(w, r, NewErrResponse(http.StatusNotFound, code, message))
}

func UnprocessableEntity(w http.ResponseWriter, r *http.Request, code, message string) {
	NewErr(w, r, NewErrResponse(http.StatusUnprocessableEntity, code, message))
}

func UnprocessableEntityIfInvalid(w http.ResponseWriter, r *http.Request, generic any, message string) bool {
	e := NewValidationErrResponse(generic, message)
	if e != nil {
		NewErr(w, r, e)
		return true
	}
	return false
}

func InternalServerError(w http.ResponseWriter, r *http.Request, code, message string) {
	NewErr(w, r, NewErrResponse(http.StatusInternalServerError, code, message))
}
