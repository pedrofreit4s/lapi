package lapi

import (
	"net/http"
	"net/url"
)

// OutOfContext retorna uma nova requisição fora do contexto.
func OutOfContext() *request {
	return &request{
		headers: make(map[string]string),
		query:   make(url.Values),
	}
}

// Send envia a requisição HTTP e retorna a resposta.
// Retorna a resposta HTTP e um erro, se ocorrer algum problema.
func (r *request) Send() (*http.Response, error) {
	req, err := http.NewRequest(r.method, r.baseURL, r.body)
	if err != nil {
		return nil, err
	}
	for key, value := range r.headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
