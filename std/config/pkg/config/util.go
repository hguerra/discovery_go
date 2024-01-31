package config

import (
	"encoding/json"
	"os"
	"strconv"
)

// loadMap loads a JSON file to map[string]any.
// Returns nil and error if not found.
func loadMap(file string) (map[string]any, error) {
	f, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var cfg map[string]any
	err = json.Unmarshal(f, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// searchMap recursively searches for a value for path in source map.
// Returns nil if not found.
// Note: This assumes that the path entries and map keys are lower cased.
// Based on: https://github.com/spf13/viper/blob/0e822151189419a986b48b3f7ffcc6fdaee29cbe/viper.go#L664
func searchMap(source map[string]any, path []string) (any, bool) {
	if len(path) == 0 {
		return nil, false
	}
	next, ok := source[path[0]]
	if ok {
		// Fast path
		if len(path) == 1 {
			return next, true
		}

		// Nested case
		switch next := next.(type) {
		case map[string]any:
			// Type assertion is safe here since it is only reached
			// if the type of `next` is the same as the type being asserted
			return searchMap(next, path[1:])
		case []interface{}:
			if len(next) == 0 {
				return nil, false
			}

			nestedArray := make(map[string]any)
			for idx, v := range next {
				nestedArray[strconv.Itoa(idx)] = v
			}

			return searchMap(nestedArray, path[1:])
		default:
			// got a value but nested key expected, return "nil" for not found
			return nil, false
		}
	}
	return nil, false
}
