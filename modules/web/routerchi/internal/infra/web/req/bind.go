package req

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/golang/gddo/httputil/header"
	"github.com/tidwall/gjson"
)

// Based on:
// https://www.alexedwards.net/blog/how-to-properly-parse-a-json-request-body
type malformedRequestError struct {
	status int
	msg    string
}

func (mr *malformedRequestError) Error() string {
	return mr.msg
}

func validateContentType(r *http.Request) error {
	// If the Content-Type header is present, check that it has the value
	// application/json. Note that we are using the gddo/httputil/header
	// package to parse and extract the value here, so the check works
	// even if the client includes additional charset or boundary
	// information in the header
	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			return &malformedRequestError{status: http.StatusUnsupportedMediaType, msg: msg}
		}
	}
	return nil
}

func enforceLimitRequestBody(w http.ResponseWriter, r *http.Request) {
	// Use http.MaxBytesReader to enforce a maximum read of 10MB from the
	// response body. A request body larger than that will now result in
	// Decode() returning a "http: request body too large" error.
	var maxBodySize int64 = 10485760
	r.Body = http.MaxBytesReader(w, r.Body, maxBodySize)
}

func Body(w http.ResponseWriter, r *http.Request) ([]byte, error) {
	err := validateContentType(r)
	if err != nil {
		return nil, err
	}
	enforceLimitRequestBody(w, r)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		if err.Error() == "http: request body too large" {
			return nil, &malformedRequestError{status: http.StatusRequestEntityTooLarge, msg: "Request body too large"}
		}
		return nil, err
	}
	if len(body) == 0 {
		return nil, &malformedRequestError{status: http.StatusRequestEntityTooLarge, msg: "Request body mandatory"}
	}
	return body, nil
}

func BindJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	err := validateContentType(r)
	if err != nil {
		return err
	}
	enforceLimitRequestBody(w, r)

	// Setup the decoder and call the DisallowUnknownFields() method on it.
	// This will cause Decode() to return a "json: unknown field ..." error
	// if it encounters any extra unexpected fields in the JSON. Strictly
	// speaking, it returns an error for "keys which do not match any
	// non-ignored, exported fields in the destination".
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err = dec.Decode(&dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError

		switch {
		// Catch any syntax errors in the JSON and send an error message
		// which interpolates the location of the problem to make it
		// easier for the client to fix.
		case errors.As(err, &syntaxError):
			msg := fmt.Sprintf("Request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
			return &malformedRequestError{status: http.StatusBadRequest, msg: msg}

		// In some circumstances Decode() may also return an
		// io.ErrUnexpectedEOF error for syntax errors in the JSON. There
		// is an open issue regarding this at
		// https://github.com/golang/go/issues/25956.
		case errors.Is(err, io.ErrUnexpectedEOF):
			msg := "Request body contains badly-formed JSON"
			return &malformedRequestError{status: http.StatusBadRequest, msg: msg}

		// Catch any type errors, like trying to assign a string in the
		// JSON request body to a int field in our Person struct. We can
		// interpolate the relevant field name and position into the error
		// message to make it easier for the client to fix.
		case errors.As(err, &unmarshalTypeError):
			msg := fmt.Sprintf(
				"Request body contains an invalid value for the %q field (at position %d)",
				unmarshalTypeError.Field,
				unmarshalTypeError.Offset,
			)
			return &malformedRequestError{status: http.StatusBadRequest, msg: msg}

		// Catch the error caused by extra unexpected fields in the request
		// body. We extract the field name from the error message and
		// interpolate it in our custom error message. There is an open
		// issue at https://github.com/golang/go/issues/29035 regarding
		// turning this into a sentinel error.
		case strings.HasPrefix(err.Error(), "json: unknown field "):
			fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
			msg := fmt.Sprintf("Request body contains unknown field %s", fieldName)
			return &malformedRequestError{status: http.StatusBadRequest, msg: msg}

		// An io.EOF error is returned by Decode() if the request body is
		// empty.
		case errors.Is(err, io.EOF):
			msg := "Request body must not be empty"
			return &malformedRequestError{status: http.StatusBadRequest, msg: msg}

		// Catch the error caused by the request body being too large. Again
		// there is an open issue regarding turning this into a sentinel
		// error at https://github.com/golang/go/issues/30715.
		case err.Error() == "http: request body too large":
			return &malformedRequestError{status: http.StatusRequestEntityTooLarge, msg: "Request body too large"}

		// Otherwise default to logging the error and sending a 500 Internal
		// Server Error response.
		default:
			return err
		}
	}

	err = dec.Decode(&struct{}{})
	if !errors.Is(err, io.EOF) {
		msg := "Request body must only contain a single JSON object"
		return &malformedRequestError{status: http.StatusBadRequest, msg: msg}
	}

	return nil
}

func BindJSONObject(w http.ResponseWriter, r *http.Request) (gjson.Result, error) {
	bytes, err := Body(w, r)
	if err != nil {
		return gjson.Result{}, err
	}
	return gjson.ParseBytes(bytes), nil
}
