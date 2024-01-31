package config

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	KEY_APP_ENV     = "APP_ENV"
	DEFAULT_APP_ENV = "development"
)

type configuration struct {
	mu       sync.RWMutex
	raw      map[string]any
	notFound map[string]bool
	strings  map[string]string
	floats64 map[string]float64
}

func (c *configuration) getConfig(key string) (any, bool) {
	path := strings.Split(key, ".")
	c.mu.RLock()
	defer c.mu.RUnlock()
	return searchMap(c.raw, path)
}

func (c *configuration) lookupEnv(key string) (any, bool) {
	path := strings.Replace(strings.ToUpper(key), ".", "_", -1)
	v, found := os.LookupEnv(path)
	if found {
		return v, true
	}
	return nil, false
}

func (c *configuration) GetEnv() string {
	c.mu.RLock()
	cv, found := c.strings[KEY_APP_ENV]
	c.mu.RUnlock()
	if found {
		return cv
	}

	e := GetEnv()
	c.mu.Lock()
	c.strings[KEY_APP_ENV] = e
	c.mu.Unlock()
	return e
}

func (c *configuration) GetString(key string) (string, bool) {
	c.mu.RLock()
	v, found := c.strings[key]
	c.mu.RUnlock()
	if found {
		return v, true
	}

	c.mu.RLock()
	_, notFound := c.notFound[key]
	c.mu.RUnlock()
	if notFound {
		return "", false
	}

	ev, found := c.lookupEnv(key)
	if found {
		cv, ok := ev.(string)
		if !ok {
			return "", false
		}

		c.mu.Lock()
		c.strings[key] = cv
		c.mu.Unlock()
		return cv, true
	}

	rv, found := c.getConfig(key)
	if found {
		cv, ok := rv.(string)
		if !ok {
			return "", false
		}

		c.mu.Lock()
		c.strings[key] = cv
		c.mu.Unlock()
		return cv, true
	}

	c.notFound[key] = true
	return "", false
}

func (c *configuration) GetFloat64(key string) (float64, bool) {
	c.mu.RLock()
	v, found := c.floats64[key]
	c.mu.RUnlock()
	if found {
		return v, true
	}

	c.mu.RLock()
	_, notFound := c.notFound[key]
	c.mu.RUnlock()
	if notFound {
		return 0, false
	}

	ev, found := c.lookupEnv(key)
	if found {
		tmpV, ok := ev.(string)
		if !ok {
			return 0, false
		}

		cv, err := strconv.ParseFloat(tmpV, 64)
		if err != nil {
			return 0, false
		}

		c.mu.Lock()
		c.floats64[key] = cv
		c.mu.Unlock()
		return cv, true
	}

	rv, found := c.getConfig(key)
	if found {
		cv, ok := rv.(float64)
		if !ok {
			return 0, false
		}

		c.mu.Lock()
		c.floats64[key] = cv
		c.mu.Unlock()
		return cv, true
	}

	c.notFound[key] = true
	return 0, false
}

func (c *configuration) GetInt8(key string) (int8, bool) {
	v, found := c.GetFloat64(key)
	if found {
		return int8(v), found
	}
	return 0, false
}

func (c *configuration) GetInt16(key string) (int16, bool) {
	v, found := c.GetFloat64(key)
	if found {
		return int16(v), found
	}
	return 0, false
}

func (c *configuration) GetInt32(key string) (int32, bool) {
	v, found := c.GetFloat64(key)
	if found {
		return int32(v), found
	}
	return 0, false
}

func (c *configuration) GetInt64(key string) (int64, bool) {
	v, found := c.GetFloat64(key)
	if found {
		return int64(v), found
	}
	return 0, false
}

func GetEnv() string {
	ev, found := os.LookupEnv(KEY_APP_ENV)
	if !found {
		ev = DEFAULT_APP_ENV
	}
	return ev
}

func Load(file string) (*configuration, error) {
	cfg, err := loadMap(file)
	if err != nil {
		return nil, err
	}
	return &configuration{
		raw:      cfg,
		notFound: make(map[string]bool),
		strings:  make(map[string]string),
		floats64: make(map[string]float64),
	}, nil
}
