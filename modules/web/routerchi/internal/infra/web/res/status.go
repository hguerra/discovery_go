package res

import (
	"net/http"
)

func OK(w http.ResponseWriter, v any) {
	JSON(w, http.StatusOK, NewResponse(v))
}

func PageOf(w http.ResponseWriter, v any, page, size, total int) {
	JSON(w, http.StatusOK, NewPageableResponse(v, page, size, total))
}

func Created(w http.ResponseWriter, v any) {
	JSON(w, http.StatusCreated, NewResponse(v))
}

func Accepted(w http.ResponseWriter, v any) {
	JSON(w, http.StatusAccepted, NewResponse(v))
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func PartialContent(w http.ResponseWriter, v any) {
	JSON(w, http.StatusPartialContent, NewResponse(v))
}

func BadRequest(w http.ResponseWriter, code, message string, errs ...string) {
	NewErr(w, NewErrResponse(http.StatusBadRequest, code, message, errs...))
}

func Unauthorized(w http.ResponseWriter, code, message string, errs ...string) {
	NewErr(w, NewErrResponse(http.StatusUnauthorized, code, message, errs...))
}

func Forbidden(w http.ResponseWriter, code, message string, errs ...string) {
	NewErr(w, NewErrResponse(http.StatusForbidden, code, message, errs...))
}

func NotFound(w http.ResponseWriter, code, message string, errs ...string) {
	NewErr(w, NewErrResponse(http.StatusNotFound, code, message, errs...))
}

func UnprocessableEntity(w http.ResponseWriter, code, message string, errs ...string) {
	NewErr(w, NewErrResponse(http.StatusUnprocessableEntity, code, message, errs...))
}

func InternalServerError(w http.ResponseWriter, code, message string, errs ...string) {
	NewErr(w, NewErrResponse(http.StatusInternalServerError, code, message, errs...))
}
