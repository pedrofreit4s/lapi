package lapi

import "encoding/json"

// SetHeaders define os cabeçalhos HTTP para a requisição.
func (r *request) SetHeaders(headers map[string]string) *request {
	r.headers = headers
	return r
}

// SetHeader define um cabeçalho HTTP para a requisição.
func (r *request) SetHeader(key, value string) *request {
	r.headers[key] = value
	return r
}

// SetHeaderJSON define um cabeçalho HTTP para a requisição como JSON.
func (r *request) SetHeaderJSON(key string, value interface{}) *request {
	jsonBody, err := json.Marshal(value)
	if err != nil {
		return r
	}
	r.headers[key] = string(jsonBody)
	return r
}
