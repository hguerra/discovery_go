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

func BadRequest(w http.ResponseWriter, r *http.Request, code, message string, errs ...string) {
	NewErr(w, r, NewErrResponse(http.StatusBadRequest, code, message, errs...))
}

func Unauthorized(w http.ResponseWriter, r *http.Request, code, message string, errs ...string) {
	NewErr(w, r, NewErrResponse(http.StatusUnauthorized, code, message, errs...))
}

func Forbidden(w http.ResponseWriter, r *http.Request, code, message string, errs ...string) {
	NewErr(w, r, NewErrResponse(http.StatusForbidden, code, message, errs...))
}

func NotFound(w http.ResponseWriter, r *http.Request, code, message string, errs ...string) {
	NewErr(w, r, NewErrResponse(http.StatusNotFound, code, message, errs...))
}

func UnprocessableEntity(w http.ResponseWriter, r *http.Request, code, message string, errs ...string) {
	NewErr(w, r, NewErrResponse(http.StatusUnprocessableEntity, code, message, errs...))
}

func InternalServerError(w http.ResponseWriter, r *http.Request, code, message string, errs ...string) {
	NewErr(w, r, NewErrResponse(http.StatusInternalServerError, code, message, errs...))
}
