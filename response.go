package godo

import (
	"encoding/json"
)

// Response returns back the http code, type of data, and the presigned url to the user.
type Response struct {
	// StatusCode is the http code that will be returned back to the user.
	StatusCode int `json:"statusCode,omitempty"`
	// Headers is the information about the type of data being returned back.
	Headers map[string]string `json:"headers,omitempty"`
	// Body will contain the presigned url to upload or download files.
	Body string `json:"body,omitempty"`
}

func NewResponseJSON() Response {
	return Response{
		Headers: map[string]string{"Content-Type": "application/json"},
	}
}

func (r *Response) AddHeader(key, value string) {
	r.Headers[key] = value
}

func (r *Response) ToMap() map[string]interface{} {
	if rJSON, err := json.Marshal(*r); err == nil {
		var output map[string]interface{}
		if err = json.Unmarshal(rJSON, &output); err == nil {
			return output
		}
	}

	return nil
}
