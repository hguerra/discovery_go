package validate

import (
	"fmt"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/hguerra/discovery_go/modules/web/routerchi/internal/infra/logging"
)

var (
	validate     *validator.Validate
	validateOnce sync.Once
)

func getValidate() *validator.Validate {
	validateOnce.Do(func() {
		validate = validator.New()
	})
	return validate
}

func Validate(generic any) []string {
	err := getValidate().Struct(generic)
	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			logging.GetLogger().Errorf("Invalid validation: %v", err)
			return nil
		}

		var errs []string
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, fmt.Sprintf("%s is %s with type %s", err.StructField(), err.Tag(), err.Type()))
		}

		return errs
	}
	return nil
}
