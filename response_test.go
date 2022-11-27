package godo_test

import (
	"net/http"
	"testing"

	"github.com/Yiannis128/godo"
)

func TestNewResponseJSON(t *testing.T) {
	runMapTestsGeneric := func(m map[string]interface{}) {
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

	r := godo.NewResponseJSON()

	m := r.ToMap()

	t.Run("Empty JSON", func(t *testing.T) {
		runMapTestsGeneric(m)

		// StatusCode
		statusCodeElement, ok := m["statusCode"]
		if !ok {
			t.Fatalf("no status code element")
		}

		// ToMap has json Unmarshal that treats ints as float64 by default.
		statusCode, ok := statusCodeElement.(float64)
		if !ok {
			t.Fatalf("statusCode is not an float64")
		}

		if statusCode != http.StatusOK {
			t.Fatalf("default statusCode element is incorrect: %f", statusCode)
		}
	})

	r.StatusCode = http.StatusAccepted
	r.Body = `{"key":"value"}`

	m = r.ToMap()

	t.Run("JSON with StatusCode and Body set", func(t *testing.T) {
		runMapTestsGeneric(m)

		// StatusCode
		statusCodeElement, ok := m["statusCode"]
		if !ok {
			t.Fatalf("no status code element")
		}

		// ToMap has json Unmarshal that treats ints as float64 by default.
		statusCode, ok := statusCodeElement.(float64)
		if !ok {
			t.Fatalf("statusCode is not an float64")
		}

		if statusCode != http.StatusAccepted {
			t.Fatalf("statusCode element is incorrect: %f", statusCode)
		}

		// Body
		body, ok := m["body"]
		if !ok {
			t.Fatalf("no body element")
		}

		if body != `{"key":"value"}` {
			t.Fatalf("body element is incorrect: %s", body)
		}

	})

}
