package lapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Model representa o contexto de uma requisição HTTP.
// Contém a requisição HTTP, informações de autenticação e configurações do ambiente.
//
// Exemplo de uso:
//
//	m := NewRequest("https://api.exemplo.com", map[string]string{
//	    "Content-Type": "application/json",
//	}, 30)
//	var response Response
//	err := m.Get("/endpoint", &response)
type model struct {
	// request representa a requisição HTTP.
	request *request

	// Auth contém as informações de autenticação.
	Auth struct {
		// Token representa o token de autenticação JWT.
		// Será automaticamente adicionado como Bearer token no header Authorization.
		Token string

		// RefreshToken representa o token de atualização JWT.
		// Pode ser usado para renovar o token de autenticação quando expirado.
		RefreshToken string
	}

	// inDevelopment indica se a requisição está em ambiente de desenvolvimento.
	// Quando true, logs adicionais podem ser exibidos para debug.
	inDevelopment bool
}

// NewRequest cria uma nova instância de Model com as configurações especificadas.
//
// Parâmetros:
//   - baseURL: URL base para todas as requisições (ex: "https://api.exemplo.com")
//   - headers: Headers HTTP padrão para todas as requisições
//   - timeout: Timeout em segundos para todas as requisições
//
// Exemplo:
//
//	m := NewRequest("https://api.exemplo.com", map[string]string{
//	    "Content-Type": "application/json",
//	}, 30)
func NewRequest(baseURL string, headers map[string]string, timeout int) *model {
	return &model{
		request: &request{
			baseURL: baseURL,
			headers: headers,
			method:  "GET",
			timeout: time.Duration(timeout) * time.Second,
			query:   make(url.Values),
		},
	}
}

// MakeRequest executa uma requisição HTTP com os parâmetros especificados.
//
// Parâmetros:
//   - method: Método HTTP (GET, POST, PUT, DELETE, etc)
//   - path: Caminho do endpoint (ex: "/users")
//   - payload: Dados a serem enviados no corpo da requisição (opcional)
//   - dest: Ponteiro para a estrutura que receberá a resposta
//
// Retorna:
//   - *httpError: Erro HTTP se a requisição falhar, nil caso contrário
func (m *model) MakeRequest(method string, path string, payload *interface{}, dest interface{}) *httpError {
	m.request.method = method
	// Parse the body
	if payload != nil {
		bodyJson, err := json.Marshal(payload)
		if err != nil {
			return m.MakeError(http.StatusInternalServerError, err.Error(), "Houve um erro interno no servidor! C: 01")
		}
		m.request.body = bytes.NewBuffer(bodyJson)
	}

	// Parse the query
	qp := ""
	if strings.Contains(m.request.baseURL, "?") {
		qp = "&"
	} else {
		qp = "?"
	}

	// Make the request
	req, err := http.NewRequest(
		m.request.method,
		fmt.Sprintf("%s%s%s%s", m.request.baseURL, path, qp, m.request.query.Encode()),
		m.request.body,
	)
	if err != nil {
		return m.MakeError(http.StatusInternalServerError, err.Error(), "Houve um erro interno no servidor! C: 02")
	}

	// Parse the headers
	for k, v := range m.request.headers {
		req.Header.Add(k, v)
	}

	// Parse the auth
	if m.Auth.Token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", m.Auth.Token))
	}

	// Perform request
	client := &http.Client{
		Timeout: m.request.timeout,
	}
	start := time.Now()
	resp, err := client.Do(req)

	responseTime := time.Since(start).Milliseconds()
	status := "(408 timeout)"

	if resp != nil {
		status = resp.Status
	}

	// Log request
	logMsg := fmt.Sprintf("[%s] %s (%s) %d ms", req.Method, req.URL, status, responseTime)
	log.Println(logMsg)

	if err != nil {
		return m.MakeError(http.StatusInternalServerError, err.Error(), "Houve um erro interno no servidor! C: 03")
	}
	defer resp.Body.Close()

	// Parse response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return m.MakeError(http.StatusInternalServerError, err.Error(), "Houve um erro interno no servidor! C: 04")
	}

	if resp.Header.Get("Content-Type") == "application/pdf" {
		*dest.(*[]byte) = body
		return nil
	}

	if err := json.Unmarshal(body, dest); err != nil {
		fmt.Println(string(body))
		log.Println(err.Error())
		return nil
	}

	// Check status code
	if resp.StatusCode >= 400 {
		return m.MakeError(resp.StatusCode, string(body), "Status code >= 400")
	}

	return nil
}

// Request retorna a instância da requisição HTTP associada ao modelo.
func (m *model) Request() *request {
	return m.request
}

// Get executa uma requisição HTTP GET.
//
// Parâmetros:
//   - path: Caminho do endpoint (ex: "/users")
//   - dest: Ponteiro para a estrutura que receberá a resposta
//
// Exemplo:
//
//	var response Response
//	err := m.Get("/users", &response)
func (m *model) Get(path string, dest interface{}) *httpError {
	r := m.MakeRequest("GET", path, nil, dest)
	return r
}

// Post executa uma requisição HTTP POST.
//
// Parâmetros:
//   - path: Caminho do endpoint (ex: "/users")
//   - payload: Dados a serem enviados no corpo da requisição
//   - dest: Ponteiro para a estrutura que receberá a resposta
//
// Exemplo:
//
//	payload := map[string]string{"name": "John"}
//	var response Response
//	err := m.Post("/users", &payload, &response)
func (m *model) Post(path string, payload *interface{}, dest interface{}) *httpError {
	r := m.MakeRequest("POST", path, payload, dest)
	return r
}

// Put executa uma requisição HTTP PUT.
//
// Parâmetros:
//   - path: Caminho do endpoint (ex: "/users/1")
//   - payload: Dados a serem enviados no corpo da requisição
//   - dest: Ponteiro para a estrutura que receberá a resposta
//
// Exemplo:
//
//	payload := map[string]string{"name": "John"}
//	var response Response
//	err := m.Put("/users/1", &payload, &response)
func (m *model) Put(path string, payload *interface{}, dest interface{}) *httpError {
	r := m.MakeRequest("PUT", path, payload, dest)
	return r
}

// Delete executa uma requisição HTTP DELETE.
//
// Parâmetros:
//   - path: Caminho do endpoint (ex: "/users/1")
//   - dest: Ponteiro para a estrutura que receberá a resposta
//
// Exemplo:
//
//	var response Response
//	err := m.Delete("/users/1", &response)
func (m *model) Delete(path string, dest interface{}) *httpError {
	r := m.MakeRequest("DELETE", path, nil, dest)
	return r
}

// Patch executa uma requisição HTTP PATCH.
//
// Parâmetros:
//   - path: Caminho do endpoint (ex: "/users/1")
//   - payload: Dados a serem enviados no corpo da requisição
//   - dest: Ponteiro para a estrutura que receberá a resposta
//
// Exemplo:
//
//	payload := map[string]string{"name": "John"}
//	var response Response
//	err := m.Patch("/users/1", &payload, &response)
func (m *model) Patch(path string, payload *interface{}, dest interface{}) *httpError {
	r := m.MakeRequest("PATCH", path, payload, dest)
	return r
}
