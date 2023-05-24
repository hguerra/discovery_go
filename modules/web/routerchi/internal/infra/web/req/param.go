package req

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func Param(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

func ParamInt(r *http.Request, key string) (int, error) {
	return strconv.Atoi(Param(r, key))
}

func ParamFloat(r *http.Request, key string) (float64, error) {
	return strconv.ParseFloat(Param(r, key), 64)
}
