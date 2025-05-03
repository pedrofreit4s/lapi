package lapi

import (
	"bytes"
	"encoding/json"
	"io"
	"net/url"
	"strings"
)

// SetBody define o corpo da requisição HTTP.
func (r *request) SetBody(body io.Reader) *request {
	r.body = body
	return r
}

// SetBodyString define o corpo da requisição HTTP como uma string.
func (r *request) SetBodyString(body string) *request {
	r.body = strings.NewReader(body)
	return r
}

// SetBodyJSON define o corpo da requisição HTTP como um JSON.
func (r *request) SetBodyJSON(body interface{}) *request {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return r
	}
	r.body = bytes.NewReader(jsonBody)
	return r
}

// SetBodyFormData define o corpo da requisição HTTP como um FormData.
func (r *request) SetBodyFormData(body map[string]string) *request {
	formData := url.Values{}
	for key, value := range body {
		formData.Add(key, value)
	}

	r.body = strings.NewReader(formData.Encode())
	return r
}
