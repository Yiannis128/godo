package godo_test

import (
	"testing"

	"github.com/Yiannis128/godo"
)

func TestNewResponseJSON(t *testing.T) {
	r := godo.NewResponseJSON()

	m := r.ToMap()

	if m == nil {
		t.Fatal("map is nil")
	}

	headersElement, ok := m["headers"]
	if !ok {
		t.Fatalf("format of map is incorrect: missing headers element: %s", m)
	}

	headers, ok := headersElement.(map[string]interface{})
	if !ok {
		t.Fatalf("headers element is not a map[string]interface{}: headers element is %s", headers)
	}

	appJSONHeaderValue, ok := headers["Content-Type"]
	if !ok {
		t.Fatalf("headers map does not contain a Content-Type key/value pair")
	}

	if appJSONHeaderValue != "application/json" {
		t.Fatalf("Content-Type value incorrect: %s", appJSONHeaderValue)
	}
}
