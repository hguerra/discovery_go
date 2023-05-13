package res

import (
	"net/http"
)

type errDTO struct {
	Status  int    `json:"status,omitempty"`
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
}

type ErrResponse struct {
	Error   errDTO   `json:"error"`
	Details []errDTO `json:"details,omitempty"`
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

func NewErr(w http.ResponseWriter, r *http.Request, err *ErrResponse) {
	JSON(w, r, err.Error.Status, err)
}
