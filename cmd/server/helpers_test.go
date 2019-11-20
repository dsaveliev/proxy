package server

import (
	"net/http"
	"net/url"
	"testing"
)

func TestValidSkipValue(t *testing.T) {
	u, _ := url.ParseRequestURI("http://test.com/recipes?skip=7")
	req := &http.Request{URL: u}
	v := parseSkipValue(req, 10)
	if v != 7 {
		t.Errorf("Skip value was incorrect, got: %d, want: %d.", v, 7)
	}
}

func TestInvalidSkipValue(t *testing.T) {
	u, _ := url.ParseRequestURI("http://test.com/recipes?skip=x")
	req := &http.Request{URL: u}
	v := parseSkipValue(req, 10)
	if v != 10 {
		t.Errorf("Skip value was incorrect, got: %d, want: %d.", v, 10)
	}

	u, _ = url.ParseRequestURI("http://test.com/recipes?skip=")
	req = &http.Request{URL: u}
	v = parseSkipValue(req, 10)
	if v != 10 {
		t.Errorf("Skip value was incorrect, got: %d, want: %d.", v, 10)
	}

	u, _ = url.ParseRequestURI("http://test.com/recipes?")
	req = &http.Request{URL: u}
	v = parseSkipValue(req, 10)
	if v != 10 {
		t.Errorf("Skip value was incorrect, got: %d, want: %d.", v, 10)
	}
}

func TestValidTopValue(t *testing.T) {
	u, _ := url.ParseRequestURI("http://test.com/recipes?top=7")
	req := &http.Request{URL: u}
	v := parseTopValue(req, 10)
	if v != 7 {
		t.Errorf("Top value was incorrect, got: %d, want: %d.", v, 7)
	}
}

func TestInvalidTopValue(t *testing.T) {
	u, _ := url.ParseRequestURI("http://test.com/recipes?top=x")
	req := &http.Request{URL: u}
	v := parseSkipValue(req, 10)
	if v != 10 {
		t.Errorf("Top value was incorrect, got: %d, want: %d.", v, 10)
	}

	u, _ = url.ParseRequestURI("http://test.com/recipes?top=")
	req = &http.Request{URL: u}
	v = parseSkipValue(req, 10)
	if v != 10 {
		t.Errorf("Top value was incorrect, got: %d, want: %d.", v, 10)
	}

	u, _ = url.ParseRequestURI("http://test.com/recipes?")
	req = &http.Request{URL: u}
	v = parseSkipValue(req, 10)
	if v != 10 {
		t.Errorf("Top value was incorrect, got: %d, want: %d.", v, 10)
	}
}

func TestValidIdsValue(t *testing.T) {
	u, _ := url.ParseRequestURI("http://test.com/recipes?ids=1,27,3")
	req := &http.Request{URL: u}
	v := parseIdsValue(req)
	expected := []int{1, 27, 3}
	for i, id := range v {
		if id != expected[i] {
			t.Errorf("Id value was incorrect, got: %d, want: %d.", id, expected[i])
		}
	}
}

func TestInvalidIdsValue(t *testing.T) {
	u, _ := url.ParseRequestURI("http://test.com/recipes?ids=1,x,3,8.")
	req := &http.Request{URL: u}
	v := parseIdsValue(req)
	expected := []int{1, 3}
	for i, id := range v {
		if id != expected[i] {
			t.Errorf("Id value was incorrect, got: %d, want: %d.", id, expected[i])
		}
	}

	u, _ = url.ParseRequestURI("http://test.com/recipes?ids=")
	req = &http.Request{URL: u}
	v = parseIdsValue(req)
	if len(v) != 0 {
		t.Errorf("Ids value was incorrect, got: %#v, want: [].", v)
	}

	u, _ = url.ParseRequestURI("http://test.com/recipes?")
	req = &http.Request{URL: u}
	v = parseIdsValue(req)
	if len(v) != 0 {
		t.Errorf("Ids value was incorrect, got: %#v, want: [].", v)
	}
}
