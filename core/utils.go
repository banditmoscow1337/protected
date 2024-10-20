package core

import (
	"io"
	"net/http"
)

func get(uri string) ([]byte, error) {
	c := &http.Client{}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}

	ret, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(ret.Body)
}

func bool2string(v bool) string {
	if v {
		return "true"
	}

	return "false"
}
