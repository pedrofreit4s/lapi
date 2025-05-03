package lapi

import (
	"net/url"
)

// SetQuery define a query string para a requisição HTTP.
func (r *request) SetQuery(query map[string]string) *request {
	queryString := url.Values{}
	for key, value := range query {
		queryString.Add(key, value)
	}

	r.query = queryString
	return r
}
