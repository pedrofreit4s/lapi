package lapi

import (
	"io"
	"net/url"
	"time"
)

// Request representa uma requisição HTTP.
// Contém todas as informações necessárias para realizar uma requisição HTTP como
// URL base, método HTTP, cabeçalhos, query parameters, corpo da requisição e timeout.
//
// Exemplo de uso:
//
//	r := NewRequest()
//	r.SetBaseURL("https://api.exemplo.com")
//	r.SetMethod("GET")
//	r.SetHeader("Content-Type", "application/json")
//	resp, err := r.Send()
type request struct {
	// BaseURL é a URL base para a requisição HTTP.
	// Exemplo: "https://api.exemplo.com"
	baseURL string

	// Method é o método HTTP a ser utilizado (GET, POST, PUT, DELETE, etc).
	// Por padrão, se não especificado, será GET.
	method string

	// Headers são os cabeçalhos HTTP a serem enviados com a requisição.
	// Exemplo: {"Content-Type": "application/json", "Authorization": "Bearer token"}
	headers map[string]string

	// Query é a query string para a requisição HTTP.
	// Exemplo: ?parametro1=valor1&parametro2=valor2
	query url.Values

	// Body é o corpo da requisição HTTP.
	// Pode ser qualquer implementação de io.Reader.
	body io.Reader

	// Timeout é o tempo máximo de resposta da requisição HTTP.
	// Se não especificado, será usado o timeout padrão do cliente HTTP.
	timeout time.Duration
}

// SetBaseURL define a URL base para a requisição HTTP.
// A URL base deve incluir o protocolo (http:// ou https://).
//
// Exemplo:
//
//	r.SetBaseURL("https://api.exemplo.com")
//
// Retorna a própria requisição para permitir encadeamento de métodos.
func (r *request) SetBaseURL(baseURL string) *request {
	r.baseURL = baseURL
	return r
}

// SetMethod define o método HTTP para a requisição.
// Os métodos suportados são: GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS.
//
// Exemplo:
//
//	r.SetMethod("POST")
//
// Retorna a própria requisição para permitir encadeamento de métodos.
func (r *request) SetMethod(method string) *request {
	r.method = method
	return r
}
