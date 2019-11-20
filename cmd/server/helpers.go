package server

import (
	"net/http"
	"strconv"
	"strings"
)

func parseTopValue(req *http.Request, defaultValue int) int {
	result := defaultValue
	if value, ok := req.URL.Query()["top"]; ok && len(value) > 0 {
		if top, err := strconv.Atoi(value[0]); err == nil {
			result = top
		}
	}
	return result
}

func parseSkipValue(req *http.Request, defaultValue int) int {
	result := defaultValue
	if value, ok := req.URL.Query()["skip"]; ok && len(value) > 0 {
		if skip, err := strconv.Atoi(value[0]); err == nil {
			result = skip
		}
	}
	return result
}

func parseIdsValue(req *http.Request) []int {
	result := []int{}
	if value, ok := req.URL.Query()["ids"]; ok && len(value) > 0 {
		vs := strings.Split(value[0], ",")
		for _, v := range vs {
			id, err := strconv.Atoi(v)
			if err != nil {
				continue
			}
			result = append(result, id)
		}
	}
	return result
}
