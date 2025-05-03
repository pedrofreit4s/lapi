package lapi

import (
	"io"
	"net/http"
	"net/url"
)

// request representa uma requisição HTTP.
// Contém informações necessárias para realizar uma requisição HTTP como
// URL base, método HTTP, cabeçalhos e corpo da requisição.
type request struct {
	// BaseURL é a URL base para a requisição HTTP.
	baseURL string
	// Method é o método HTTP a ser utilizado (GET, POST, PUT, DELETE, etc).
	method string
	// Headers são os cabeçalhos HTTP a serem enviados com a requisição.
	headers map[string]string
	// Query é a query string para a requisição HTTP.
	query url.Values
	// Body é o corpo da requisição HTTP.
	body io.Reader
	// Dest é o destino para a resposta HTTP.
	dest interface{}
}

// SetBaseURL define a URL base para a requisição HTTP.
// Retorna a própria requisição para permitir encadeamento de métodos.
func (r *request) SetBaseURL(baseURL string) *request {
	r.baseURL = baseURL
	return r
}

// SetMethod define o método HTTP para a requisição.
// Retorna a própria requisição para permitir encadeamento de métodos.
func (r *request) SetMethod(method string) *request {
	r.method = method
	return r
}

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
