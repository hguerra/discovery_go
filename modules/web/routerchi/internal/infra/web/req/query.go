package req

import (
	"net/http"
	"strconv"
)

func QueryHas(r *http.Request, key string) bool {
	return r.URL.Query().Has(key)
}

func Query(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func QueryAll(r *http.Request, key string) []string {
	if QueryHas(r, key) {
		return r.URL.Query()[key]
	}
	return make([]string, 0)
}

func DefaultQuery(r *http.Request, key, defaultValue string) string {
	value := Query(r, key)
	if value == "" {
		return defaultValue
	}
	return value
}

func QueryInt(r *http.Request, key string) (int, error) {
	return strconv.Atoi(Query(r, key))
}

func DefaultQueryInt(r *http.Request, key string, defaultValue int) int {
	value, err := QueryInt(r, key)
	if err != nil {
		return defaultValue
	}
	return value
}

func QueryFloat(r *http.Request, key string) (float64, error) {
	return strconv.ParseFloat(Query(r, key), 64)
}

func DefaultQueryFloat(r *http.Request, key string, defaultValue float64) float64 {
	value, err := QueryFloat(r, key)
	if err != nil {
		return defaultValue
	}
	return value
}
