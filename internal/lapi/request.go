package lapi

import (
	"io"
	"net/url"
	"time"
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
	// Timeout é o tempo máximo de resposta da requisição HTTP.
	timeout time.Duration
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
