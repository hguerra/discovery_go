package res

import (
	"net/http"

	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/web/validate"
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

func NewErrResponse(status int, code, message string) *ErrResponse {
	return &ErrResponse{
		Error: errDTO{
			Status:  status,
			Code:    code,
			Message: message,
		},
	}
}

func NewValidationErrResponse(generic any, message string) *ErrResponse {
	errs := validate.Validate(generic)
	if len(errs) == 0 {
		return nil
	}

	var details []errDTO
	for _, e := range errs {
		details = append(details, errDTO{
			Message: e,
		})
	}

	return &ErrResponse{
		Error: errDTO{
			Status:  http.StatusUnprocessableEntity,
			Message: message,
		},
		Details: details,
	}
}

func NewErr(w http.ResponseWriter, r *http.Request, err *ErrResponse) {
	JSON(w, r, err.Error.Status, err)
}
