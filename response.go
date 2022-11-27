package godo

import (
	"encoding/json"
	"net/http"
)

// Response - The object that represents the response that will be sent back to
// Digital Ocean serverless function / manager. Calling ToMap() will return the
// ResponseMap that will apply the interface.
type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

// NewResponseJSON - Creates a Response struct with the application/json header
// Content-Type already added.
func NewResponseJSON() Response {
	return Response{
		Headers: map[string]string{"Content-Type": "application/json"},
	}
}

// AddHeader - Adds a value to the header.
func (r *Response) AddHeader(key, value string) {
	r.Headers[key] = value
}

// ToMap - Returns a ResponseMap object that can be returned from the Main
// function of a serverless function package.
func (r *Response) ToMap() ResponseMap {
	response := *r
	// There is no 0 status code so give a default value.
	if response.StatusCode == 0 {
		response.StatusCode = http.StatusOK
	}

	if rJSON, err := json.Marshal(response); err == nil {
		var output map[string]interface{}
		if err = json.Unmarshal(rJSON, &output); err == nil {
			return output
		}
	}

	return nil
}

// ResponseMap can be returned from the Main function of a serverless function.
type ResponseMap map[string]interface{}
